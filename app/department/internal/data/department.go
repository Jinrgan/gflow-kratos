package data

import (
	"context"
	"gflow-kratos/app/department/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type departmentRepo struct {
	data *Data
	log  *log.Helper
}

// NewDepartmentRepo .
func NewDepartmentRepo(data *Data, logger log.Logger) biz.DepartmentRepo {
	return &departmentRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *departmentRepo) CreateGreeter(ctx context.Context, g *biz.Department) error {
	return nil
}

func (r *departmentRepo) UpdateGreeter(ctx context.Context, g *biz.Department) error {
	return nil
}
