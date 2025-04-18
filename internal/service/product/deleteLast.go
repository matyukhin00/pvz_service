package product

import "context"

func (s *ProductService) DeleteLast(ctx context.Context, receptionId string) error {
	return s.repository.DeleteLast(ctx, receptionId)
}
