package pvz

import (
	"context"
	"errors"
	"fmt"

	sq "github.com/Masterminds/squirrel"

	"github.com/matyukhin00/pvz_service/internal/model"
)

func (r *PvzRepository) Create(ctx context.Context, info model.Pvz) (*model.Pvz, error) {
	builder := sq.Insert(table).
		PlaceholderFormat(sq.Dollar).
		Columns(columnId, columnCity, columnRegistrationDate).
		Values(info.Id, info.City, info.RegistrationDate).
		Suffix(fmt.Sprintf("RETURNING %s, %s, %s", columnId, columnCity, columnRegistrationDate))

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	ans := &model.Pvz{}
	err = r.pool.QueryRow(ctx, query, args...).Scan(&ans.Id, &ans.City, &ans.RegistrationDate)
	if err != nil {
		return nil, errors.New("Pvz with that ID exists")
	}

	return ans, nil
}
