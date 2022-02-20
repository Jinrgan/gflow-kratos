package shift

import (
	"context"
	"errors"
	pb "gflow-kratos/api/shift/v1"
	pkgMysql "gflow-kratos/pkg/mysql"
	"strconv"

	"github.com/go-kratos/kratos/v2/log"

	"github.com/go-sql-driver/mysql"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Repo interface {
	CreateAffair(ctx context.Context, aff *pb.Affair) error
	GetShiftBySchedule(ctx context.Context, schId int32) (*pb.Shift, error)
	GetShifts(ctx context.Context) ([]*pb.Shift, error)
}

type Service struct {
	//pb.UnimplementedHandoverServiceServer

	//UAAClt *uaaSwagger.APIClient
	//SchClt *schSwagger.APIClient
	Repo Repo
}

func (h *Service) CreateAffair(ctx context.Context, aff *pb.Affair) (*emptypb.Empty, error) {
	err := h.Repo.CreateAffair(ctx, aff)
	if errors.Is(err, strconv.ErrSyntax) {
		return nil, status.Errorf(codes.InvalidArgument, "affair %+v to create is invalid", aff)
	}
	var sqlErr *mysql.MySQLError
	if errors.As(err, &sqlErr) && sqlErr.Number == pkgMysql.ErDupKey {
		return nil, status.Error(codes.AlreadyExists, "affair already exist")
	}
	if err != nil {
		log.Errorf("cannot create affair: %v", err)
		return nil, status.Error(codes.Unknown, "")
	}

	return nil, nil
}

func (h *Service) GetAffairByCurrentUser(ctx context.Context, _ *emptypb.Empty) (*pb.Affair, error) {
	panic("implement me")
}

func (h *Service) GetAffairs(ctx context.Context, _ *emptypb.Empty) (*pb.GetAffairsResponse, error) {
	panic("implement me")
}

func (h *Service) CreateShift(ctx context.Context, req *pb.CreateShiftRequest) (*emptypb.Empty, error) {
	panic("implement me")
}

func (h *Service) GetShifts(ctx context.Context, _ *emptypb.Empty) (*pb.GetShiftsResponse, error) {
	panic("implement me")
}
