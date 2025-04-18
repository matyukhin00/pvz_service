package service

import (
	"context"

	"github.com/matyukhin00/pvz_service/internal/model"
)

type UserService interface {
	DummyLogin(ctx context.Context, info model.UserClaims) (string, error)
	Create(ctx context.Context, info model.User) (*model.User, error)
	Login(ctx context.Context, info model.User) (string, error)
	ValidateToken(tokenStr string) (*model.UserClaims, error)
	//Get()
	//Update()
	//Delete()
}

type PvzService interface {
	Create(ctx context.Context, info model.Pvz) (*model.Pvz, error)
	Exists(ctx context.Context, id string) (bool, error)
}

type ReceptionService interface {
	Create(ctx context.Context, info string) (*model.Reception, error)
	Close(ctx context.Context, info string) (*model.Reception, error)
	GetId(ctx context.Context, pvzId string) (string, error)
}

type ProductService interface {
	Add(ctx context.Context, info model.AddProduct) (*model.Product, error)
}
