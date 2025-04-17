package reception

import "context"

func (r *ReceptionRepository) ExistsOpen(ctx context.Context, info string) (bool, error) {
	var result bool
	err := r.pool.QueryRow(ctx, "SELECT EXISTS (SELECT 1 FROM receptions WHERE pvz_id = $1 AND status = 'in_progress')", info).Scan(&result)
	if err != nil {
		return false, err
	}

	return result, nil
}
