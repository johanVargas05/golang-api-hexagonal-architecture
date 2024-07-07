package manager_cache_services

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/johanVargas05/golang-api-hexagonal-architecture/mocks"
	dtos "github.com/johanVargas05/golang-api-hexagonal-architecture/src/domain/dtos/manager_cache"
)

func TestNewManagerCacheService(t *testing.T) {
	mockRepository := mocks.NewManagerCacheRepositoryPort(t)

	service := New(mockRepository)

	assert.NotNil(t, service)
	assert.Equal(t, mockRepository, service.repository)
}

func TestGetData(t *testing.T) {
	mockRepository := mocks.NewManagerCacheRepositoryPort(t)
	service := New(mockRepository)

	key := "testKey"
	structure := struct {
		Field string
	}{}

	params := dtos.CacheParams{
		Key:   key,
		Value: structure,
	}

	t.Run("returns data when repository returns data", func(t *testing.T) {
		mockRepository.On("GetData", mock.Anything, &params).Return(nil)

		result, err := service.GetData(key, structure)

		mockRepository.ExpectedCalls = nil

		assert.Nil(t, err)
		assert.Equal(t, structure, result)
	})

	t.Run("returns error when repository returns error", func(t *testing.T) {
		mockRepository.On("GetData", mock.Anything, &params).Return(errors.New("test error"))
		_, err := service.GetData(key, structure)

		mockRepository.ExpectedCalls = nil

		assert.NotNil(t, err)
		assert.Equal(t, "test error", err.Error())
	})
}

func TestSetData(t *testing.T) {
	mockRepository := mocks.NewManagerCacheRepositoryPort(t)
	service := New(mockRepository)

	ctx := context.Background()
	key := "testKey"
	structure := struct {
		Field string
	}{}

	params := dtos.CacheParams{
		Key:   key,
		Value: structure,
	}
	expiration := time.Minute

	t.Run("does not register error log when repository does not return error", func(t *testing.T) {
		mockRepository.On("SetData", ctx, &params, expiration).Return(nil)

		service.SetData(key, structure, expiration)

		mockRepository.ExpectedCalls = nil

		mockRepository.AssertExpectations(t)
	})

	t.Run("registers error log when repository returns error", func(t *testing.T) {
		mockRepository.On("SetData", ctx, &params, expiration).Return(errors.New("test error"))

		service.SetData(key, structure, expiration)

		mockRepository.AssertExpectations(t)
	})
}
