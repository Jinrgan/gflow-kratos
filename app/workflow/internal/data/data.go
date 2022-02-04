package data

import (
	"gflow-kratos/app/workflow/internal/conf"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
	"gorm.io/gorm/schema"

	kratosLog "github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewGormDB, NewData, NewWorkflowRepo, NewProcessRepo)

// Data .
type Data struct {
	db *gorm.DB
}

func NewGormDB(conf *conf.Data, logger kratosLog.Logger) *gorm.DB {
	hlp := kratosLog.NewHelper(kratosLog.With(logger, "module", "workflow-service/data/gorm"))

	db, err := gorm.Open(mysql.Open(conf.GetDatabase().Source), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 启用单数命名
		},
		Logger: gormLogger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			gormLogger.Config{
				SlowThreshold: time.Millisecond, // 慢查询阈值
				Colorful:      true,
				LogLevel:      gormLogger.Info,
			}),
	})
	if err != nil {
		hlp.Fatalf("failed opening connection to db: %v", err)
	}

	db.AutoMigrate()

	return db
}

// NewData .
func NewData(db *gorm.DB, logger kratosLog.Logger) (*Data, func(), error) {
	hlp := kratosLog.NewHelper(kratosLog.With(logger, "module", "workflow-service/data"))

	return &Data{db: db}, func() {
		hlp.Info("closing the data resources")

		sqlDB, err := db.DB()
		if err != nil {
			hlp.Error(err)
		}

		err = sqlDB.Close()
		if err != nil {
			hlp.Error(err)
		}
	}, nil
}
