package user

import (
	"context"
	"time"

	"github.com/matyukhin00/pvz_service/internal/model"
)

func (s *UserService) DummyLogin(ctx context.Context, info model.UserClaims) (string, error) {
	token, err := s.token.GenerateToken(
		info,
		[]byte(secretKey),
		time.Hour*24,
	)
	if err != nil {
		return "", err
	}

	return token, nil
}
