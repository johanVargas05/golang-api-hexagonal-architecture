package seed_ports

type RunSeedServicePort interface {
	Execute() error
}

type SeedRunCheckServicePort interface {
	Execute() error
}
