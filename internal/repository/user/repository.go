package user

import (
	"github.com/jackc/pgx/v4/pgxpool"
)

const (
	table             = "users"
	columnId          = "id"
	columnEmail       = "email"
	columnPassword    = "password_hash"
	columnRole        = "role"
	columnCreatedTime = "created_at"
)

type UserRepository struct {
	pool *pgxpool.Pool
}

func NewUserRepository(pool *pgxpool.Pool) *UserRepository {
	return &UserRepository{
		pool: pool,
	}
}
