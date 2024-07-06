package run_seed_service

import seed_ports "github.com/johanVargas05/golang-api-hexagonal-architecture/src/domain/ports/seed"

type RunSeedService struct {
	runSeedRepository seed_ports.RunSeedRepositoryPort
	loadDataSeedRepository seed_ports.LoadDataSeedRepositoryPort
}

func New(runSeedRepository seed_ports.RunSeedRepositoryPort, loadDataSeedRepository seed_ports.LoadDataSeedRepositoryPort) *RunSeedService {
	return &RunSeedService{
		runSeedRepository,
		loadDataSeedRepository,
	}
}

func (service *RunSeedService) Execute() error {
	portfolio, err := service.loadDataSeedRepository.Execute()
	
	if err != nil {
		return err
	}

	err = service.runSeedRepository.Execute(portfolio)
	
	if err != nil {
		return err
	}

	return nil
}
