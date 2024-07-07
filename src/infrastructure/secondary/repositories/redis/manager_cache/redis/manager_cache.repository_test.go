package manager_cache_repository

import (
	"errors"
	"testing"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/go-redis/redismock/v8"
	"github.com/stretchr/testify/assert"

	dtos "github.com/johanVargas05/golang-api-hexagonal-architecture/src/domain/dtos/manager_cache"
)

func TestNewManagerCacheRepository(t *testing.T) {
	db, _ := redismock.NewClientMock()

	repository := New(db)

	assert.NotNil(t, repository)
}

func TestGetData(t *testing.T) {
	db, mockClient := redismock.NewClientMock()
	repository := New(db)

	key := "testKey"
	structure := struct {
		Test string
	}{}
	params := dtos.CacheParams{
		Key:   key,
		Value: structure,
	}

	t.Run("success", func(t *testing.T) {

		value := `{"Test": "value"}`
		mockClient.ExpectGet(key).SetVal(value)

		err := repository.GetData(&params)

		assert.Nil(t, err)
		assert.Equal(t, map[string]interface{}{"Test": "value"}, params.Value)
	})

	t.Run("redis.Nil error", func(t *testing.T) {
		mockClient.ExpectGet(key).SetErr(redis.Nil)

		err := repository.GetData(&params)

		assert.NotNil(t, err)
	})

	t.Run("redis error get data", func(t *testing.T) {
		errorMock := errors.New("error mock")
		mockClient.ExpectGet(key).SetErr(errorMock)

		err := repository.GetData(&params)

		assert.Equal(t, errorMock, err)
	})

	t.Run("json.Unmarshal error", func(t *testing.T) {
		value := `{"test":`
		mockClient.ExpectGet(key).SetVal(value)

		err := repository.GetData(&params)

		assert.NotNil(t, err)
	})
}

func TestSetData(t *testing.T) {
	db, mockClient := redismock.NewClientMock()
	repository := New(db)

	key := "testKey"
	valueMap := struct {
		Test string
	}{}
	params := dtos.CacheParams{
		Key:   key,
		Value: valueMap,
	}
	expiration := time.Minute

	t.Run("success", func(t *testing.T) {
		mockClient.ExpectSet(key, `{"Test":""}`, expiration).SetVal("OK")

		err := repository.SetData(&params, expiration)

		assert.Nil(t, err)
		assert.Nil(t, mockClient.ExpectationsWereMet())
	})

	t.Run("json.Marshal error", func(t *testing.T) {
		invalidValue := make(chan int)
		params.Value = invalidValue

		err := repository.SetData(&params, expiration)

		assert.NotNil(t, err)
	})

	t.Run("redis error set data", func(t *testing.T) {
		errorMock := errors.New("redis error")
		mockClient.ExpectSet(key, `{"Test":"value"}`, expiration).SetErr(errorMock)
		valueMap := struct {
			Test string
		}{
			Test: "value",
		}
		params.Value = valueMap

		err := repository.SetData(&params, expiration)

		assert.Equal(t, errorMock, err)
	})
}
