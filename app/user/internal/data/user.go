package data

import (
	"context"
	"gflow-kratos/app/user/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type userRepo struct {
	data *Data
	log  *log.Helper
}

// NewUserRepo .
func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *userRepo) CreateGreeter(ctx context.Context, g *biz.User) error {
	return nil
}

func (r *userRepo) UpdateGreeter(ctx context.Context, g *biz.User) error {
	return nil
}
