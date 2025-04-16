package user

import (
	"context"
	"net/mail"

	"github.com/matyukhin00/pvz_service/internal/model"
	"github.com/matyukhin00/pvz_service/internal/repository"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repository repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{
		repository: repo,
	}
}

func (s *UserService) Create(ctx context.Context, info model.UserRequest) (*model.UserAnswer, error) {
	_, err := mail.ParseAddress(info.Email)
	if err != nil {
		return nil, errors.New("Invalid email address")
	}

	if len(info.Password) < 4 {
		return nil, errors.New("Your password must be 4 or more symbols")
	}

	if info.Role != "employee" && info.Role != "moderator" {
		return nil, errors.New("Your role must be `employee` or `moderator`")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(info.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	info.Password = string(hashedPassword)

	return s.repository.Create(ctx, info)
}
