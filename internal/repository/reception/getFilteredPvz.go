package reception

import (
	"context"
	"fmt"
	"time"
)

func (r *ReceptionRepository) GetFilteredPvz(ctx context.Context, start, end time.Time) ([]string, error) {
	rows, err := r.pool.Query(
		ctx,
		fmt.Sprintf("SELECT DISTINCT %s FROM %s WHERE %s BETWEEN $1 AND $2", columnPvzId, table, columnDate),
		start,
		end,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result := make([]string, 0)
	for rows.Next() {
		var id string
		err = rows.Scan(&id)
		if err != nil {
			return nil, err
		}

		result = append(result, id)
	}

	return result, nil
}
