package repository

import (
	"context"
	"latihan_grpc/modules/model"
)

type RepoUser interface {
	Show(ctx context.Context) (userRes []model.User, err error)
	Create(ctx context.Context, userIn model.User) (userRes model.User, err error)
	FindById(ctx context.Context, userId uint64) (userRes model.User, err error)
	Update(ctx context.Context, userIn model.User) (userRes model.User, err error)
	Delete(ctx context.Context, userId uint64) (err error)
}
