package reception

import "github.com/matyukhin00/pvz_service/internal/repository"

type ReceptionService struct {
	repository repository.ReceptionRepository
}

func NewReceptionService(repo repository.ReceptionRepository) *ReceptionService {
	return &ReceptionService{
		repository: repo,
	}
}
