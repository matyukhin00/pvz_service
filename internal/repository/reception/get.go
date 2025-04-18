package reception

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/matyukhin00/pvz_service/internal/model"
)

func (r *ReceptionRepository) Get(ctx context.Context, id string) (*model.Reception, error) {
	builder := sq.Select("*").
		From(table).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{columnId: id})

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	res := &model.Reception{}
	err = r.pool.QueryRow(ctx, query, args...).Scan(&res.Id, &res.PvzId, &res.DateTime, &res.Status)
	if err != nil {
		return nil, err
	}

	return res, nil
}
