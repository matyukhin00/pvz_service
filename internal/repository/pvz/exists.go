package pvz

import (
	"context"
)

func (r *PvzRepository) Exists(ctx context.Context, id string) (bool, error) {
	var result bool
	err := r.pool.QueryRow(ctx, "SELECT EXISTS (SELECT 1 FROM pvz WHERE id = $1)", id).Scan(&result)
	if err != nil {
		return false, err
	}

	return result, nil
}
