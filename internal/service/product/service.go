package product

import "github.com/matyukhin00/pvz_service/internal/repository"

type ProductService struct {
	repository repository.ProductRepository
}

func NewProductService(repo repository.ProductRepository) *ProductService {
	return &ProductService{
		repository: repo,
	}
}
