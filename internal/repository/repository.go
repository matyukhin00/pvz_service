package repository

import (
	"context"

	"github.com/matyukhin00/pvz_service/internal/model"
)

type UserRepository interface {
	Create(ctx context.Context, info model.UserRequest) (*model.UserAnswer, error)
	//Read(id int) (any, error)
	//Update(info any) (any, error)
	//Delete(info any) (int, error)
}
