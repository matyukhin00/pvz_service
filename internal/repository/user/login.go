package user

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/matyukhin00/pvz_service/internal/model"
	"github.com/pkg/errors"
)

func (r *UserRepository) Login(ctx context.Context, info model.User) (*model.User, error) {
	builder := sq.Select(columnId, columnPassword, columnRole).
		From(table).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{columnEmail: info.Email}).
		Limit(1)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, errors.New("Failed to build query")
	}

	req := &model.User{}

	err = r.pool.QueryRow(ctx, query, args...).Scan(&req.Id, &req.Password, &req.Role)
	if err != nil {
		return nil, errors.New("User with that email does not exist")
	}

	return req, nil
}
