package service

import (
	"context"
	"time"

	"github.com/matyukhin00/pvz_service/internal/model"
)

type UserService interface {
	DummyLogin(ctx context.Context, info model.UserClaims) (string, error)
	Create(ctx context.Context, info model.User) (*model.User, error)
	Login(ctx context.Context, info model.User) (string, error)
	ValidateToken(tokenStr string) (*model.UserClaims, error)
}

type PvzService interface {
	Create(ctx context.Context, info model.Pvz) (*model.Pvz, error)
	Exists(ctx context.Context, id string) (bool, error)
	Get(ctx context.Context, id string) (*model.Pvz, error)
}

type ReceptionService interface {
	Create(ctx context.Context, info string) (*model.Reception, error)
	Close(ctx context.Context, info string) (*model.Reception, error)
	GetId(ctx context.Context, pvzId string) (string, error)
	Get(ctx context.Context, id string) (*model.Reception, error)
	GetFilteredPvz(ctx context.Context, start, end time.Time) ([]string, error)
	GetAll(ctx context.Context, pvzId string) ([]model.Reception, error)
}

type ProductService interface {
	Add(ctx context.Context, info model.AddProduct) (*model.Product, error)
	DeleteLast(ctx context.Context, receptionId string) error
	GetAll(ctx context.Context, receptionId string) ([]model.Product, error)
}
