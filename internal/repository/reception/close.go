package reception

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/matyukhin00/pvz_service/internal/model"
)

func (r *ReceptionRepository) Close(ctx context.Context, info string) (*model.Reception, error) {
	builder := sq.Update(table).
		Set(columnStatus, "close").
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{columnPvzId: info}).
		Suffix(fmt.Sprintf("RETURNING %s, %s, %s, %s", columnId, columnDate, columnPvzId, columnStatus))

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	req := &model.Reception{}

	err = r.pool.QueryRow(ctx, query, args...).Scan(&req.Id, &req.DateTime, &req.PvzId, &req.Status)
	if err != nil {
		return nil, err
	}

	return req, nil
}
