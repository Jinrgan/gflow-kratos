package service

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	pb "gflow-kratos/api/department/v1"
	"gflow-kratos/app/department/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

// DepartmentService is a greeter service.
type DepartmentService struct {
	pb.UnimplementedDepartmentServer

	uc  *biz.DepartmentUseCase
	log *log.Helper
}

// NewDepartmentService new a department service.
func NewDepartmentService(uc *biz.DepartmentUseCase, logger log.Logger) *DepartmentService {
	return &DepartmentService{uc: uc, log: log.NewHelper(logger)}
}

func (s *DepartmentService) GetDepartments(ctx context.Context, _ *emptypb.Empty) (*pb.GetDepartmentsResponse, error) {
	panic("implement me")
}
