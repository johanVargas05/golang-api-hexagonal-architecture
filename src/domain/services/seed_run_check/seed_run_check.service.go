package seed_run_check_service

import seed_ports "github.com/johanVargas05/golang-api-hexagonal-architecture/src/domain/ports/seed"

type SeedRunCheckService struct {
	seedRunCheckRepository seed_ports.SeedRunCheckRepositoryPort
}

func New(seedRunCheckRepository seed_ports.SeedRunCheckRepositoryPort) *SeedRunCheckService {
	return &SeedRunCheckService{
		seedRunCheckRepository,
	}
}

func (service *SeedRunCheckService) Execute() error {
	err := service.seedRunCheckRepository.Execute()

	if err != nil {
		return err
	}

	return nil
}