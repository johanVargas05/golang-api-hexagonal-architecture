// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	entities "github.com/johanVargas05/golang-api-hexagonal-architecture/src/domain/entities"
	mock "github.com/stretchr/testify/mock"

	portfolio_dtos "github.com/johanVargas05/golang-api-hexagonal-architecture/src/domain/dtos/portfolio"
)

// AllPortFolioOfUserQueryRepositoryPort is an autogenerated mock type for the AllPortFolioOfUserQueryRepositoryPort type
type AllPortFolioOfUserQueryRepositoryPort struct {
	mock.Mock
}

// Execute provides a mock function with given fields: params
func (_m *AllPortFolioOfUserQueryRepositoryPort) Execute(params *portfolio_dtos.ParamsGetAllPortfolioOfUserDto) ([]*entities.Portfolio, int, error) {
	ret := _m.Called(params)

	if len(ret) == 0 {
		panic("no return value specified for Execute")
	}

	var r0 []*entities.Portfolio
	var r1 int
	var r2 error
	if rf, ok := ret.Get(0).(func(*portfolio_dtos.ParamsGetAllPortfolioOfUserDto) ([]*entities.Portfolio, int, error)); ok {
		return rf(params)
	}
	if rf, ok := ret.Get(0).(func(*portfolio_dtos.ParamsGetAllPortfolioOfUserDto) []*entities.Portfolio); ok {
		r0 = rf(params)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*entities.Portfolio)
		}
	}

	if rf, ok := ret.Get(1).(func(*portfolio_dtos.ParamsGetAllPortfolioOfUserDto) int); ok {
		r1 = rf(params)
	} else {
		r1 = ret.Get(1).(int)
	}

	if rf, ok := ret.Get(2).(func(*portfolio_dtos.ParamsGetAllPortfolioOfUserDto) error); ok {
		r2 = rf(params)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// NewAllPortFolioOfUserQueryRepositoryPort creates a new instance of AllPortFolioOfUserQueryRepositoryPort. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewAllPortFolioOfUserQueryRepositoryPort(t interface {
	mock.TestingT
	Cleanup(func())
}) *AllPortFolioOfUserQueryRepositoryPort {
	mock := &AllPortFolioOfUserQueryRepositoryPort{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
