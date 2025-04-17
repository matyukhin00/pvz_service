package reception

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/matyukhin00/pvz_service/internal/model"
)

func (r *ReceptionRepository) Create(ctx context.Context, info string) (*model.Reception, error) {

	builder := sq.Insert(table).
		PlaceholderFormat(sq.Dollar).
		Columns(columnPvzId).
		Values(info).
		Suffix(fmt.Sprintf("RETURNING %s, %s, %s, %s", columnId, columnPvzId, columnDate, columnStatus))

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	ans := &model.Reception{}
	err = r.pool.QueryRow(ctx, query, args...).Scan(&ans.Id, &ans.PvzId, &ans.DateTime, &ans.Status)
	if err != nil {
		return nil, err
	}

	return ans, nil
}
