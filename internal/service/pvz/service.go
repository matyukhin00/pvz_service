package pvz

import (
	"github.com/matyukhin00/pvz_service/internal/repository"
)

type PvzService struct {
	repository repository.PvzRepository
}

func NewPvzService(repo repository.PvzRepository) *PvzService {
	return &PvzService{
		repository: repo,
	}
}
