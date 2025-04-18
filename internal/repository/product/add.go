package product

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/matyukhin00/pvz_service/internal/model"
)

func (r *ProductRepository) Add(ctx context.Context, info model.AddProduct) (*model.Product, error) {
	builder := sq.Insert(table).
		PlaceholderFormat(sq.Dollar).
		Columns(columnReceptionId, columnType).
		Values(info.ReceptionId, info.Type).
		Suffix(fmt.Sprintf(
			"RETURNING %s, %s, %s, %s",
			columnId,
			columnDate,
			columnType,
			columnReceptionId,
		))

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	result := &model.Product{}
	err = r.pool.QueryRow(ctx, query, args...).Scan(
		&result.Id,
		&result.DateTime,
		&result.Type,
		&result.ReceptionId,
	)

	if err != nil {
		return nil, err
	}

	return result, nil
}
