package service

import (
	"context"

	"github.com/matyukhin00/pvz_service/internal/model"
)

type UserService interface {
	Create(ctx context.Context, info model.UserRequest) (*model.UserAnswer, error)
	//Get()
	//Update()
	//Delete()
}
