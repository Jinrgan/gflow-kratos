package shift

import (
	"context"
	pb "gflow-kratos/api/shift/v1"
	pkgMysql "gflow-kratos/pkg/mysql"
	"strconv"
	"testing"

	"github.com/go-sql-driver/mysql"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestHandler_CreateAffair(t *testing.T) {
	var svc Service
	re := &repo{}
	svc.Repo = re

	cases := []struct {
		name        string
		repoErr     error
		wantErrCode codes.Code
	}{
		{
			name: "normal_affair",
		},
		{
			name:        "bad_request",
			repoErr:     strconv.ErrSyntax,
			wantErrCode: codes.InvalidArgument,
		},
		{
			name: "already_exist",
			repoErr: &mysql.MySQLError{
				Number: pkgMysql.ErDupKey,
			},
			wantErrCode: codes.AlreadyExists,
		},
	}

	for _, cc := range cases {
		t.Run(cc.name, func(t *testing.T) {
			re.err = cc.repoErr
			_, err := svc.CreateAffair(context.Background(), &pb.Affair{})
			if err != nil && status.Code(err) == cc.wantErrCode || err == nil {
				return
			} else {
				t.Errorf("want status code: %d; got: %d", cc.wantErrCode, status.Code(err))
			}
		})
	}
}

type repo struct {
	shift *pb.Shift
	err   error
}

func (r *repo) CreateAffair(ctx context.Context, aff *pb.Affair) error {
	return r.err
}

func (r *repo) GetShiftBySchedule(ctx context.Context, schId int32) (*pb.Shift, error) {
	if r.err != nil {
		return nil, r.err
	}

	return r.shift, nil
}

func (r *repo) GetShifts(ctx context.Context) ([]*pb.Shift, error) {
	panic("implement me")
}
