package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

type User struct {
	Hello string
}

type UserRepo interface {
	CreateGreeter(context.Context, *User) error
	UpdateGreeter(context.Context, *User) error
}

type UserUseCase struct {
	repo UserRepo
	log  *log.Helper
}

func NewUserUseCase(repo UserRepo, logger log.Logger) *UserUseCase {
	return &UserUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *UserUseCase) Create(ctx context.Context, g *User) error {
	return uc.repo.CreateGreeter(ctx, g)
}

func (uc *UserUseCase) Update(ctx context.Context, g *User) error {
	return uc.repo.UpdateGreeter(ctx, g)
}
