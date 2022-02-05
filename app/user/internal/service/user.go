package service

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	pb "gflow-kratos/api/user/v1"
	"gflow-kratos/app/user/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

// UserService is a greeter service.
type UserService struct {
	pb.UnimplementedUserServer

	uc  *biz.UserUseCase
	log *log.Helper
}

// NewUserService new a user service.
func NewUserService(uc *biz.UserUseCase, logger log.Logger) *UserService {
	return &UserService{uc: uc, log: log.NewHelper(logger)}
}

func (s *UserService) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	panic("implement me")
}

func (s *UserService) GetUsers(ctx context.Context, _ *emptypb.Empty) (*pb.GetUsersResponse, error) {
	panic("implement me")
}
