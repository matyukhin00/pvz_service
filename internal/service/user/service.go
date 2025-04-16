package user

import (
	"os"

	"github.com/matyukhin00/pvz_service/internal/repository"
)

var secretKey string

type UserService struct {
	repository repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{
		repository: repo,
	}
}

func init() {
	secretKey = os.Getenv("secretKey")
}
