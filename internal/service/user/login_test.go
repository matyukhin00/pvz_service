package user

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/joho/godotenv"
	"github.com/matyukhin00/pvz_service/internal/model"
	"github.com/matyukhin00/pvz_service/internal/repository"
	"github.com/matyukhin00/pvz_service/internal/repository/mocks"
	"github.com/matyukhin00/pvz_service/internal/utils"
	mocksG "github.com/matyukhin00/pvz_service/internal/utils/mocks"
	"golang.org/x/crypto/bcrypt"
)

func TestUserService_Login(t *testing.T) {
	godotenv.Load(".env")
	secretKey = os.Getenv("secretKey")

	mockRepo := mocks.NewUserRepositoryMock(t)
	mockGen := mocksG.NewTokenGeneratorMock(t)

	password := "Password12345678"
	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 4)

	type fields struct {
		repository repository.UserRepository
		token      utils.TokenGenerator
	}
	type args struct {
		ctx  context.Context
		info model.User
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "invalid email format",
			fields: fields{
				repository: mockRepo,
				token:      mockGen,
			},
			args: args{
				ctx: context.Background(),
				info: model.User{
					Email:    "bad_email",
					Password: "password",
				},
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "Ivalid password",
			fields: fields{
				repository: mockRepo,
				token:      mockGen,
			},
			args: args{
				ctx: context.Background(),
				info: model.User{
					Email:    "user@example.com",
					Password: "password",
				},
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "Wrong password",
			fields: fields{
				repository: mockRepo,
				token:      mockGen,
			},
			args: args{
				ctx: context.Background(),
				info: model.User{
					Email:    "user@example.com",
					Password: "password",
				},
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "Correct password",
			fields: fields{
				repository: mockRepo,
				token:      mockGen,
			},
			args: args{
				ctx: context.Background(),
				info: model.User{
					Email:    "user@example.com",
					Password: password,
					Role:     "employee",
				},
			},
			want:    "token",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &UserService{
				repository: tt.fields.repository,
				token:      mockGen,
			}

			if tt.name == "Wrong password" || tt.name == "Correct password" {
				mockRepo.LoginMock.Expect(tt.args.ctx, tt.args.info).Return(&model.User{
					Password: string(hashPassword),
					Role:     tt.args.info.Role,
				}, nil)
			} else {
				mockRepo.LoginMock.Expect(tt.args.ctx, tt.args.info).Return(&model.User{
					Email:    tt.args.info.Email,
					Password: tt.args.info.Password,
					Role:     "employee",
				}, nil)
			}

			mockGen.GenerateTokenMock.Expect(model.UserClaims{
				Role: "employee",
			}, []byte(secretKey), time.Hour*24).Return("token", nil)

			got, err := s.Login(tt.args.ctx, tt.args.info)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserService.Login() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("UserService.Login() = %v, want %v", got, tt.want)
			}

		})
	}
}
