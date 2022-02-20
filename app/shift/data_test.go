package shift

import (
	"context"
	"errors"
	pb "gflow-kratos/api/shift/v1"
	pkgMysql "gflow-kratos/pkg/mysql"
	mysqlTesting "gflow-kratos/pkg/mysql/testing"
	"os"
	"testing"

	"github.com/go-sql-driver/mysql"
)

var mysqlURL = "handover?charset=utf8mb4&parseTime=True&loc=Local"

func TestZeroString(t *testing.T) {
	db, err := mysqlTesting.NewDefaultDB(mysqlURL)
	if err != nil {
		t.Fatalf("cannot connect mysql: %v", err)
	}

	err = db.AutoMigrate(&Affair{})
	if err != nil {
		t.Fatalf("fail to auto migrate: %v", err)
	}

	m := Mysql{DB: db}

	err = m.CreateAffair(context.Background(), &pb.Affair{
		MciScheduleId: "1",
	})
	if err != nil {
		t.Errorf("cannot create affair: %v", err)
	}
}

func TestMysql_CreateAffair(t *testing.T) {
	db, err := mysqlTesting.NewDB()
	if err != nil {
		t.Fatalf("cannot connect mysql: %v", err)
	}

	err = db.AutoMigrate(&Affair{})
	if err != nil {
		t.Fatalf("fail to auto migrate: %v", err)
	}

	m := Mysql{DB: db}

	cases := []struct {
		name        string
		schId       string
		planId      string
		wantErrCode uint16
	}{
		{
			name:   "one_affair",
			schId:  "233",
			planId: "233",
		},
		{
			name:   "another_affair",
			schId:  "333",
			planId: "333",
		},
		{
			name:        "affair_exist",
			schId:       "233",
			planId:      "233",
			wantErrCode: pkgMysql.ErDupKey,
		},
		//{
		//	name:        "affair_without_plan",
		//	schId:       "433",
		//	planId:      "0",
		//	wantErrCode: 2,
		//},
	}

	for _, cc := range cases {
		t.Run(cc.name, func(t *testing.T) {
			err = m.CreateAffair(context.Background(), &pb.Affair{
				MciScheduleId: cc.schId,
				PlanType:      pb.PlanType_PLAN_TYPE_PATROL_PLAN,
				PlanId:        cc.planId,
			})
			if cc.wantErrCode != 0 {
				var sqlErr *mysql.MySQLError
				if err == nil {
					t.Error("want err; got none")
				} else if !errors.As(err, &sqlErr) {
					t.Errorf("want sql error; got unknown error")
				} else if sqlErr.Number != cc.wantErrCode {
					t.Errorf("want error code: %d; got: %d", cc.wantErrCode, sqlErr.Number)
				} else {
					return
				}
			}
			if err != nil {
				t.Errorf("operation failed: %v", err)
			}
		})
	}
}

func TestMain(m *testing.M) {
	os.Exit(mysqlTesting.RunWithMysqlInDocker(m))
}
