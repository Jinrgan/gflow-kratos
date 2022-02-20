package main

import (
	"fmt"
	pb "gflow-kratos/api/shift/v1"
	"gflow-kratos/app/shift"
	"time"

	"github.com/go-kratos/kratos/contrib/registry/consul/v2"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/go-kratos/swagger-api/openapiv2"
	consulAPI "github.com/hashicorp/consul/api"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	traceSDK "go.opentelemetry.io/otel/sdk/trace"
	semConv "go.opentelemetry.io/otel/semconv/v1.4.0"
	gormMysql "gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
	"gorm.io/gorm/schema"

	stdLog "log"
	"os"
)

const mysqlURL = "root:123456@tcp(localhost:3306)/shift?charset=utf8mb4&parseTime=True&loc=Local"

// go build -ldflags "-X main.Version=x.y.z"
var (
	// Name is the name of the compiled software.
	Name = "gflow.shift"
	// Version is the version of the compiled software.
	Version string
)

func main() {
	logger := log.With(log.NewStdLogger(os.Stdout),
		"service.name", Name,
		"service.version", Version,
		"ts", log.DefaultTimestamp,
		"caller", log.DefaultCaller,
	)

	db, err := gorm.Open(
		gormMysql.Open(fmt.Sprintf(mysqlURL)),
		&gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true, // 启用单数命名
			},
			Logger: gormLogger.New(
				stdLog.New(os.Stdout, "\r\n", stdLog.LstdFlags),
				gormLogger.Config{
					SlowThreshold: time.Second, // 慢查询阈值
					Colorful:      true,
					LogLevel:      gormLogger.Info,
				}),
		})
	if err != nil {
		log.Fatalf("cannot connect to database: %v", err)
	}
	err = db.AutoMigrate(&shift.Affair{})
	if err != nil {
		log.Fatalf("fail to auto migrate: %v", err)
	}

	exp, err := jaeger.New(jaeger.WithCollectorEndpoint())
	if err != nil {
		panic(err)
	}
	tp := traceSDK.NewTracerProvider(
		traceSDK.WithBatcher(exp),
		traceSDK.WithResource(resource.NewSchemaless(
			semConv.ServiceNameKey.String(Name),
		)),
	)

	regtrar, err := NewRegistrar()
	if err != nil {
		log.Fatalf("cannot create registrar: %v", err)
	}

	log.Fatal(kratos.New(
		kratos.Name(Name),
		kratos.Version(Version),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(logger),
		kratos.Server(NewServer(&shift.Config{}, &shift.Service{Repo: &shift.Mysql{DB: db}}, tp)),
		kratos.Registrar(regtrar),
	).Run())
}

func NewServer(conf *shift.Config, svc *shift.Service, tp *traceSDK.TracerProvider) *http.Server {
	opts := []http.ServerOption{
		http.Middleware(
			tracing.Server(tracing.WithTracerProvider(tp)),
			logging.Server(log.NewStdLogger(os.Stdout)),
		),
	}
	if conf.Network != "" {
		opts = append(opts, http.Network(conf.Network))
	}
	if conf.Addr != "" {
		opts = append(opts, http.Address(conf.Addr))
	}
	if conf.Timeout != nil {
		opts = append(opts, http.Timeout(*conf.Timeout))
	}

	srv := http.NewServer(opts...)
	srv.HandlePrefix("/openapi/", openapiv2.NewHandler())

	pb.RegisterShiftServiceHTTPServer(srv, svc)

	return srv
}

func NewRegistrar() (*consul.Registry, error) {
	consulAPI.DefaultConfig()
	clt, err := consulAPI.NewClient(consulAPI.DefaultConfig())
	if err != nil {
		return nil, fmt.Errorf("cannot create consul client: %w", err)
	}

	return consul.New(clt, consul.WithHealthCheck(false)), nil
}
