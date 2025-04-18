package pvz

import (
	"context"

	"github.com/matyukhin00/pvz_service/internal/model"
)

func (s *PvzService) Get(ctx context.Context, id string) (*model.Pvz, error) {
	return s.repository.Get(ctx, id)
}
