package service

import (
	"context"

	"github.com/matyukhin00/pvz_service/internal/model"
)

type UserService interface {
	DummyLogin(ctx context.Context, info model.UserClaims) (string, error)
	Create(ctx context.Context, info model.User) (*model.User, error)
	Login(ctx context.Context, info model.User) (string, error)
	//Get()
	//Update()
	//Delete()
}
