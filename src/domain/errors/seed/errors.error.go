package seed_errors

import "errors"

var (
	ErrSeedAlreadyExecuted = errors.New("seed already executed")
	ErrSeedNotExecuted     = errors.New("seed not executed")
	ErrLoadDataSeed        = errors.New("error loading data seed")
)

