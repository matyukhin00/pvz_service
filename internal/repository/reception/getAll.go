package reception

import (
	"context"
	"fmt"

	"github.com/matyukhin00/pvz_service/internal/model"
)

func (r *ReceptionRepository) GetAll(ctx context.Context, pvzId string) ([]model.Reception, error) {
	rows, err := r.pool.Query(
		ctx,
		fmt.Sprintf("SELECT * FROM %s WHERE %s = $1", table, columnPvzId),
		pvzId,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result := make([]model.Reception, 0)
	for rows.Next() {
		rec := &model.Reception{}
		err = rows.Scan(&rec.Id, &rec.PvzId, &rec.DateTime, &rec.Status)
		if err != nil {
			return nil, err
		}

		result = append(result, *rec)
	}

	return result, nil
}
