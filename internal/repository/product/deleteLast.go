package product

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
)

func (r *ProductRepository) DeleteLast(ctx context.Context, receptionId string) error {
	req, err := r.pool.Exec(
		ctx,
		fmt.Sprintf(
			"DELETE FROM %s WHERE %s = (SELECT %s FROM %s WHERE %s = $1 ORDER BY %s DESC LIMIT 1)",
			table,
			columnId,
			columnId,
			table,
			columnReceptionId,
			columnDate,
		),
		receptionId,
	)
	if err != nil {
		return err
	}
	if req.RowsAffected() == 0 {
		return errors.New("There is no products in reception")
	}
	return nil
}
