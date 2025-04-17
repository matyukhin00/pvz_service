package pvz

import "context"

func (s *PvzService) Exists(ctx context.Context, id string) (bool, error) {
	return s.repository.Exists(ctx, id)
}
