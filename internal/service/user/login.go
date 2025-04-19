package user

import (
	"context"
	"net/mail"
	"time"

	"github.com/matyukhin00/pvz_service/internal/model"
	"github.com/pkg/errors"

	"github.com/matyukhin00/pvz_service/internal/utils"
)

func (s *UserService) Login(ctx context.Context, info model.User) (string, error) {
	_, err := mail.ParseAddress(info.Email)
	if err != nil {
		return "", errors.New("Invalid email")
	}

	if len(info.Password) < 4 {
		return "", errors.New("Invalid password")
	}

	ans, err := s.repository.Login(ctx, info)
	if err != nil {
		return "", err
	}

	if !utils.VerifyPassword(ans.Password, info.Password) {
		return "", errors.New("Invalid password")
	}

	token, err := s.token.GenerateToken(
		model.UserClaims{
			Id:   ans.Id,
			Role: ans.Role,
		},
		[]byte(secretKey),
		time.Hour*24,
	)

	return token, nil
}
