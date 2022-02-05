package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

type Department struct {
	Hello string
}

type DepartmentRepo interface {
	CreateGreeter(context.Context, *Department) error
	UpdateGreeter(context.Context, *Department) error
}

type DepartmentUseCase struct {
	repo DepartmentRepo
	log  *log.Helper
}

func NewDepartmentUseCase(repo DepartmentRepo, logger log.Logger) *DepartmentUseCase {
	return &DepartmentUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *DepartmentUseCase) Create(ctx context.Context, g *Department) error {
	return uc.repo.CreateGreeter(ctx, g)
}

func (uc *DepartmentUseCase) Update(ctx context.Context, g *Department) error {
	return uc.repo.UpdateGreeter(ctx, g)
}
