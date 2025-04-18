package product

import (
	"context"

	"github.com/matyukhin00/pvz_service/internal/model"
	"github.com/pkg/errors"
)

func (s *ProductService) Add(ctx context.Context, info model.AddProduct) (*model.Product, error) {
	if info.Type != "электроника" && info.Type != "одежда" && info.Type != "обувь" {
		return nil, errors.New("Type must be 'электроника', 'одежда' or 'обувь'")
	}

	return s.repository.Add(ctx, info)
}
