package product

import (
	"context"

	"github.com/matyukhin00/pvz_service/internal/model"
)

func (s *ProductService) GetAll(ctx context.Context, receptionId string) ([]model.Product, error) {
	return s.repository.GetAll(ctx, receptionId)
}
