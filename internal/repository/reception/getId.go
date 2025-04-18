package reception

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
)

func (r *ReceptionRepository) GetId(ctx context.Context, pvzId string) (string, error) {
	var id string
	err := r.pool.QueryRow(
		ctx,
		fmt.Sprintf("SELECT %s FROM %s WHERE %s = $1 AND %s = 'in_progress'", columnId, table, columnPvzId, columnStatus),
		pvzId,
	).Scan(&id)

	if err != nil {
		return "", errors.New("There is no open reception in that PVZ")
	}

	return id, nil
}
