package run_seed_use_case

import (
	seed_ports "github.com/johanVargas05/golang-api-hexagonal-architecture/src/domain/ports/seed"
)

type RunSeedUseCase struct {
	runSeedService seed_ports.RunSeedServicePort
	seedRunCheckService seed_ports.SeedRunCheckServicePort
}

func New(runSeedService seed_ports.RunSeedServicePort, seedRunCheckService seed_ports.SeedRunCheckServicePort) RunSeedUseCase {
	return RunSeedUseCase{
		runSeedService,
		seedRunCheckService,
	}
}

func (useCase RunSeedUseCase) Execute() error {
	err := useCase.seedRunCheckService.Execute()
	if err != nil {
		return err
	}

	err = useCase.runSeedService.Execute()
	if err != nil {
		return err
	}

	return nil
}