package reception

import (
	"context"

	"github.com/matyukhin00/pvz_service/internal/model"
)

func (s *ReceptionService) GetAll(ctx context.Context, pvzId string) ([]model.Reception, error) {
	return s.repository.GetAll(ctx, pvzId)
}
