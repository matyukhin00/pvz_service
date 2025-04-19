package user

import (
	"context"
	"reflect"
	"testing"

	"github.com/matyukhin00/pvz_service/internal/model"
	"github.com/matyukhin00/pvz_service/internal/repository"
	"github.com/matyukhin00/pvz_service/internal/repository/mocks"
)

func TestUserService_Create(t *testing.T) {
	mockRepo := mocks.NewUserRepositoryMock(t)

	type fields struct {
		repository repository.UserRepository
	}
	type args struct {
		ctx  context.Context
		info model.User
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.User
		wantErr bool
	}{
		{
			name: "invalid email format",
			fields: fields{
				repository: mockRepo,
			},
			args: args{
				ctx: context.Background(),
				info: model.User{
					Email:    "bad_email",
					Password: "password",
					Role:     "employee",
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Ivalid password",
			fields: fields{
				repository: mockRepo,
			},
			args: args{
				ctx: context.Background(),
				info: model.User{
					Email:    "user@example.com",
					Password: "pas",
					Role:     "employee",
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Invalid role",
			fields: fields{
				repository: mockRepo,
			},
			args: args{
				ctx: context.Background(),
				info: model.User{
					Email:    "user@example.com",
					Password: "password",
					Role:     "role",
				},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &UserService{
				repository: tt.fields.repository,
			}
			got, err := s.Create(tt.args.ctx, tt.args.info)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserService.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserService.Create() = %v, want %v", got, tt.want)
			}
		})
	}
}
