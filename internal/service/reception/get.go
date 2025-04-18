package reception

import (
	"context"

	"github.com/matyukhin00/pvz_service/internal/model"
)

func (s *ReceptionService) Get(ctx context.Context, id string) (*model.Reception, error) {
	return s.repository.Get(ctx, id)
}
