package pvz

import (
	"context"
	"fmt"

	"github.com/matyukhin00/pvz_service/internal/model"
)

func (r *PvzRepository) Get(ctx context.Context, id string) (*model.Pvz, error) {
	res := &model.Pvz{}
	err := r.pool.QueryRow(
		ctx,
		fmt.Sprintf("SELECT * FROM %s WHERE %s = $1", table, columnId),
		id,
	).Scan(&res.Id, &res.City, &res.RegistrationDate)

	if err != nil {
		return nil, err
	}

	return res, nil
}
