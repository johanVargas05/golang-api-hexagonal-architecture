package run_seed_use_case

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/johanVargas05/golang-api-hexagonal-architecture/mocks"
	seed_errors "github.com/johanVargas05/golang-api-hexagonal-architecture/src/domain/errors/seed"
)

func TestNew(t *testing.T) {
	runSeedServiceMock:= mocks.NewRunSeedServicePort(t)
	seedRunCheckServiceMock:= mocks.NewSeedRunCheckServicePort(t)
	useCase:= New(runSeedServiceMock, seedRunCheckServiceMock)

	assert.NotNil(t, useCase)
	assert.Equal(t, runSeedServiceMock, useCase.runSeedService)
	assert.Equal(t, seedRunCheckServiceMock, useCase.seedRunCheckService)
}

func TestExecute(t *testing.T) {
	runSeedServiceMock:= mocks.NewRunSeedServicePort(t)
	seedRunCheckServiceMock:= mocks.NewSeedRunCheckServicePort(t)
	useCase:= New(runSeedServiceMock, seedRunCheckServiceMock)

	t.Run("The seed has already been executed ❌", func(t *testing.T) {
		seedRunCheckServiceMock.On("Execute").Return(seed_errors.ErrSeedAlreadyExecuted).Once()
		
		err:= useCase.Execute()

		assert.Equal(t, seed_errors.ErrSeedAlreadyExecuted, err)
	})

	t.Run("Failed to check if the seed has already been executed ❌", func(t *testing.T) {
		errorMock:=errors.New("Error checking if the seed has already been executed")

		seedRunCheckServiceMock.On("Execute").Return(errorMock).Once()
		
		err:= useCase.Execute()

		assert.Equal(t, errorMock, err)
	})

	t.Run("Failed to run the seed ❌", func(t *testing.T) {
		seedRunCheckServiceMock.On("Execute").Return(nil).Once()
		runSeedServiceMock.On("Execute").Return(seed_errors.ErrSeedNotExecuted).Once()
		
		err:= useCase.Execute()

		assert.NotNil(t, err)
		assert.Equal(t, seed_errors.ErrSeedNotExecuted, err)
	})

	t.Run("Successfully run the seed ✅", func(t *testing.T) {
		seedRunCheckServiceMock.On("Execute").Return(nil).Once()
		runSeedServiceMock.On("Execute").Return(nil).Once()
		
		err:= useCase.Execute()

		assert.Nil(t, err)
	})

}