package product

import (
	"context"
	"fmt"

	"github.com/matyukhin00/pvz_service/internal/model"
)

func (r *ProductRepository) GetAll(ctx context.Context, receptionId string) ([]model.Product, error) {
	rows, err := r.pool.Query(
		ctx,
		fmt.Sprintf("SELECT * FROM %s WHERE %s = $1", table, columnReceptionId),
		receptionId,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result := make([]model.Product, 0)
	for rows.Next() {
		rec := &model.Product{}
		err = rows.Scan(&rec.Id, &rec.ReceptionId, &rec.DateTime, &rec.Type)
		if err != nil {
			return nil, err
		}

		result = append(result, *rec)
	}

	return result, nil
}
