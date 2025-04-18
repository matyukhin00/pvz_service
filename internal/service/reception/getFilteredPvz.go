package reception

import (
	"context"
	"time"
)

func (s *ReceptionService) GetFilteredPvz(ctx context.Context, start, end time.Time) ([]string, error) {
	return s.repository.GetFilteredPvz(ctx, start, end)
}
