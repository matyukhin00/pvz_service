package reception

import (
	"context"

	"github.com/matyukhin00/pvz_service/internal/model"
	"github.com/pkg/errors"
)

func (s *ReceptionService) Create(ctx context.Context, info string) (*model.Reception, error) {
	exists, err := s.repository.ExistsOpen(ctx, info)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, errors.New("PVZ has an open reception")
	}

	return s.repository.Create(ctx, info)
}
