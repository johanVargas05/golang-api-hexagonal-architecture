// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	entities "github.com/johanVargas05/golang-api-hexagonal-architecture/src/domain/entities"
	mock "github.com/stretchr/testify/mock"

	pagination_dtos "github.com/johanVargas05/golang-api-hexagonal-architecture/src/domain/dtos/pagination"

	portfolio_dtos "github.com/johanVargas05/golang-api-hexagonal-architecture/src/domain/dtos/portfolio"
)

// AllPortFolioOfUserCacheServicePort is an autogenerated mock type for the AllPortFolioOfUserCacheServicePort type
type AllPortFolioOfUserCacheServicePort struct {
	mock.Mock
}

// Get provides a mock function with given fields: params
func (_m *AllPortFolioOfUserCacheServicePort) Get(params *portfolio_dtos.ParamsGetAllPortfolioOfUserDto) ([]*entities.Portfolio, *pagination_dtos.PaginationResponseDto) {
	ret := _m.Called(params)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 []*entities.Portfolio
	var r1 *pagination_dtos.PaginationResponseDto
	if rf, ok := ret.Get(0).(func(*portfolio_dtos.ParamsGetAllPortfolioOfUserDto) ([]*entities.Portfolio, *pagination_dtos.PaginationResponseDto)); ok {
		return rf(params)
	}
	if rf, ok := ret.Get(0).(func(*portfolio_dtos.ParamsGetAllPortfolioOfUserDto) []*entities.Portfolio); ok {
		r0 = rf(params)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*entities.Portfolio)
		}
	}

	if rf, ok := ret.Get(1).(func(*portfolio_dtos.ParamsGetAllPortfolioOfUserDto) *pagination_dtos.PaginationResponseDto); ok {
		r1 = rf(params)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*pagination_dtos.PaginationResponseDto)
		}
	}

	return r0, r1
}

// Set provides a mock function with given fields: params, data, pagination
func (_m *AllPortFolioOfUserCacheServicePort) Set(params *portfolio_dtos.ParamsGetAllPortfolioOfUserDto, data []*entities.Portfolio, pagination *pagination_dtos.PaginationResponseDto) {
	_m.Called(params, data, pagination)
}

// NewAllPortFolioOfUserCacheServicePort creates a new instance of AllPortFolioOfUserCacheServicePort. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewAllPortFolioOfUserCacheServicePort(t interface {
	mock.TestingT
	Cleanup(func())
}) *AllPortFolioOfUserCacheServicePort {
	mock := &AllPortFolioOfUserCacheServicePort{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
