package user

import (
	"context"

	sq "github.com/Masterminds/squirrel"

	"github.com/matyukhin00/pvz_service/internal/model"
	"github.com/pkg/errors"
)

func (r *UserRepository) Create(ctx context.Context, info model.User) (*model.User, error) {
	builder := sq.Insert(table).
		PlaceholderFormat(sq.Dollar).
		Columns(columnEmail, columnPassword, columnRole).
		Values(info.Email, info.Password, info.Role).
		Suffix("RETURNING id, email, role")

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	ans := &model.User{}
	err = r.pool.QueryRow(ctx, query, args...).Scan(&ans.Id, &ans.Email, &ans.Role)
	if err != nil {
		return nil, errors.New("User with that email exists")
	}

	return ans, nil
}
