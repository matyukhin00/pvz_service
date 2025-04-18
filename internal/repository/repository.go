package repository

import (
	"context"

	"github.com/matyukhin00/pvz_service/internal/model"
)

type UserRepository interface {
	Create(ctx context.Context, info model.User) (*model.User, error)
	Login(ctx context.Context, info model.User) (*model.User, error)
}

type PvzRepository interface {
	Create(ctx context.Context, info model.Pvz) (*model.Pvz, error)
	Exists(ctx context.Context, id string) (bool, error)
}

type ReceptionRepository interface {
	Create(ctx context.Context, info string) (*model.Reception, error)
	ExistsOpen(ctx context.Context, info string) (bool, error)
	Close(ctx context.Context, info string) (*model.Reception, error)
	GetId(ctx context.Context, pvzId string) (string, error)
	Get(ctx context.Context, id string) (*model.Reception, error)
}

type ProductRepository interface {
	Add(ctx context.Context, info model.AddProduct) (*model.Product, error)
	DeleteLast(ctx context.Context, receptionId string) error
}
