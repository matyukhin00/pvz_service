package user

import (
	"os"

	"github.com/matyukhin00/pvz_service/internal/repository"
	"github.com/matyukhin00/pvz_service/internal/utils"
)

var secretKey string

type UserService struct {
	repository repository.UserRepository
	token      utils.TokenGenerator
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{
		repository: repo,
		token:      utils.NewTokenGen(),
	}
}

func init() {
	secretKey = os.Getenv("secretKey")
}
