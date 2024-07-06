package seed_ports

type RunSeedUseCasePort interface {
	Execute() error
}
