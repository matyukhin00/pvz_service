package user

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/matyukhin00/pvz_service/internal/model"
	"github.com/pkg/errors"
)

func (s *UserService) ValidateToken(tokenStr string) (*model.UserClaims, error) {
	token, err := jwt.ParseWithClaims(
		tokenStr,
		&model.UserClaims{},
		func(token *jwt.Token) (interface{}, error) {
			_, ok := token.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				return nil, errors.New("unexpected token signing method")
			}

			return []byte(secretKey), nil
		},
	)

	if err != nil {
		return nil, errors.Errorf("invalid token: %s", err.Error())
	}

	claims, ok := token.Claims.(*model.UserClaims)
	if !ok {
		return nil, errors.Errorf("invalid token claims")
	}

	//fmt.Println(tokenStr, token, claims)

	return claims, nil
}
