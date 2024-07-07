package pagination_service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPaginationService(t *testing.T) {
	t.Run("Test with valid parameters", func(t *testing.T) {
		service := New(1,10)

		assert.NotNil(t, service)
		assert.Equal(t, 1, service.page)
		assert.Equal(t, 10, service.pageSize)
	})

	t.Run("Test with zero values", func(t *testing.T) {
		service := New(0,0)

		assert.NotNil(t, service)
		assert.Equal(t, 0, service.page)
		assert.Equal(t, 0, service.pageSize)
	})
}

func TestGetPagination(t *testing.T) {
	service := New(1, 10)
	

	t.Run("Test with total items less than limit ✅", func(t *testing.T) {
		totalItems := 5
		response := service.Execute(totalItems)

		assert.Equal(t, service.page, response.CurrentPage)
		assert.Equal(t, service.pageSize, response.PageSize)
		assert.False(t, response.HasNextPage)
		assert.False(t, response.HasPreviousPage)
	})

	t.Run("Test with total items more than limit ✅", func(t *testing.T) {
		totalItems := 15
		response := service.Execute(totalItems)

		assert.Equal(t, service.page, response.CurrentPage)
		assert.Equal(t, service.pageSize, response.PageSize)
		assert.True(t, response.HasNextPage)
		assert.False(t, response.HasPreviousPage)
	})

	t.Run("Test with total items equal to limit ✅", func(t *testing.T) {
		totalItems := 10
		response := service.Execute(totalItems)

		assert.Equal(t, service.page, response.CurrentPage)
		assert.Equal(t, service.pageSize, response.PageSize)
		assert.False(t, response.HasNextPage)
		assert.False(t, response.HasPreviousPage)
	})
}

func TestGetOffset(t *testing.T) {
	t.Run("Test with page less than or equal to 0", func(t *testing.T) {
		service :=New(0,10)
		offset := service.GetOffset()
		assert.Equal(t, 0, offset)
	})

	t.Run("Test with page greater than 0", func(t *testing.T) {
		service := New(2,10)
		offset := service.GetOffset()
		assert.Equal(t, 10, offset)
	})
}

func TestGetLimit(t *testing.T) {
	t.Run("Test with limit less than or equal to 10", func(t *testing.T) {
		service := New(1,5)
		limit := service.GetLimit()
		assert.Equal(t, 10, limit)
	})

	t.Run("Test with limit greater than 10", func(t *testing.T) {
		service := New(1,25)
		limit := service.GetLimit()
		assert.Equal(t, 25, limit)
	})
}
