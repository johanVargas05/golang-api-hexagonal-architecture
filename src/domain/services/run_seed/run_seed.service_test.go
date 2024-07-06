package run_seed_service

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/johanVargas05/golang-api-hexagonal-architecture/mocks"
	"github.com/johanVargas05/golang-api-hexagonal-architecture/src/domain/entities"
	seed_errors "github.com/johanVargas05/golang-api-hexagonal-architecture/src/domain/errors/seed"
)

func TestNew(t *testing.T) {
runSeedRepositoryMock:=mocks.NewRunSeedRepositoryPort(t)
loadDataSeedRepositoryMock:=mocks.NewLoadDataSeedRepositoryPort(t)
service:=New(runSeedRepositoryMock, loadDataSeedRepositoryMock)

assert.NotNil(t, service)
assert.Equal(t, runSeedRepositoryMock, service.runSeedRepository)
assert.Equal(t, loadDataSeedRepositoryMock, service.loadDataSeedRepository)
}

func TestExecute(t *testing.T) {
	runSeedRepositoryMock:=mocks.NewRunSeedRepositoryPort(t)
	loadDataSeedRepositoryMock:=mocks.NewLoadDataSeedRepositoryPort(t)
	service:=New(runSeedRepositoryMock, loadDataSeedRepositoryMock)

	t.Run("Error loading data ❌", func(t *testing.T) {
		loadDataSeedRepositoryMock.On("Execute").Return(nil, seed_errors.ErrLoadDataSeed).Once()
		err:=service.Execute()

		assert.NotNil(t, err)
		assert.Equal(t, seed_errors.ErrLoadDataSeed, err)
	})

	t.Run("Error running seed ❌", func(t *testing.T) {
		loadDataSeedRepositoryMock.On("Execute").Return([]*entities.Portfolio{}, nil).Once()
		runSeedRepositoryMock.On("Execute", []*entities.Portfolio{}).Return(seed_errors.ErrSeedNotExecuted).Once()
		err:=service.Execute()

		assert.NotNil(t, err)
		assert.Equal(t, seed_errors.ErrSeedNotExecuted, err)
	})

	t.Run("Success ✅", func(t *testing.T) {
		loadDataSeedRepositoryMock.On("Execute").Return([]*entities.Portfolio{}, nil).Once()
		runSeedRepositoryMock.On("Execute", []*entities.Portfolio{}).Return(nil).Once()
		err:=service.Execute()

		assert.Nil(t, err)
	})
	
}