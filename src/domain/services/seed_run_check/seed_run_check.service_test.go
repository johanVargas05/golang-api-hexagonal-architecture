package seed_run_check_service

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/johanVargas05/golang-api-hexagonal-architecture/mocks"
	seed_errors "github.com/johanVargas05/golang-api-hexagonal-architecture/src/domain/errors/seed"
)

func TestNew(t *testing.T) {
	repositoryMock:=mocks.NewSeedRunCheckRepositoryPort(t)
	service:=New(repositoryMock)
	
	assert.NotNil(t, service)
	assert.Equal(t, repositoryMock, service.seedRunCheckRepository)
}

func TestExecute(t *testing.T) {
	repositoryMock:=mocks.NewSeedRunCheckRepositoryPort(t)
	service:=New(repositoryMock)
	t.Run("Error check run ❌", func(t *testing.T) {
		repositoryMock.On("Execute").Return(seed_errors.ErrSeedAlreadyExecuted).Once()
		err:=service.Execute()

		assert.NotNil(t, err)
		assert.Equal(t, seed_errors.ErrSeedAlreadyExecuted, err)
	})

	t.Run("Success ✅", func(t *testing.T) {
		repositoryMock.On("Execute").Return(nil).Once()
		err:=service.Execute()

		assert.Nil(t, err)
	})
}