package reception

import "context"

func (s *ReceptionService) GetId(ctx context.Context, pvzId string) (string, error) {
	return s.repository.GetId(ctx, pvzId)
}
