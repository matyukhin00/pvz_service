package pvz

import (
	"context"

	"github.com/matyukhin00/pvz_service/internal/model"
	"github.com/pkg/errors"
)

func (s *PvzService) Create(ctx context.Context, info model.Pvz) (*model.Pvz, error) {
	if info.City != "Москва" && info.City != "Санкт-Петербург" && info.City != "Казань" {
		return nil, errors.New("Impossible to create pvz in this city")
	}

	return s.repository.Create(ctx, info)
}
