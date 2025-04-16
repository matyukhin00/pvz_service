package service

import (
	"context"

	"github.com/matyukhin00/pvz_service/internal/model"
)

type UserService interface {
	DummyLogin(ctx context.Context, info model.UserClaims) (string, error)
	Create(ctx context.Context, info model.UserRequest) (*model.UserAnswer, error)
	Login(ctx context.Context, info model.UserLogin) (string, error)
	//Get()
	//Update()
	//Delete()
}
