// Code generated by mockery v2.14.1. DO NOT EDIT.

package mocks

import (
	teachers "mini-project/businesses/teachers"

	mock "github.com/stretchr/testify/mock"
)

// Usecase is an autogenerated mock type for the Usecase type
type Usecase struct {
	mock.Mock
}

// CreateTeacher provides a mock function with given fields: teacherDomain
func (_m *Usecase) CreateTeacher(teacherDomain *teachers.Domain) teachers.Domain {
	ret := _m.Called(teacherDomain)

	var r0 teachers.Domain
	if rf, ok := ret.Get(0).(func(*teachers.Domain) teachers.Domain); ok {
		r0 = rf(teacherDomain)
	} else {
		r0 = ret.Get(0).(teachers.Domain)
	}

	return r0
}

// Delete provides a mock function with given fields: id
func (_m *Usecase) Delete(id string) bool {
	ret := _m.Called(id)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// GetByID provides a mock function with given fields: id
func (_m *Usecase) GetByID(id string) teachers.Domain {
	ret := _m.Called(id)

	var r0 teachers.Domain
	if rf, ok := ret.Get(0).(func(string) teachers.Domain); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(teachers.Domain)
	}

	return r0
}

// GetTeachers provides a mock function with given fields:
func (_m *Usecase) GetTeachers() []teachers.Domain {
	ret := _m.Called()

	var r0 []teachers.Domain
	if rf, ok := ret.Get(0).(func() []teachers.Domain); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]teachers.Domain)
		}
	}

	return r0
}

// Login provides a mock function with given fields: userDomain
func (_m *Usecase) Login(userDomain *teachers.Domain) string {
	ret := _m.Called(userDomain)

	var r0 string
	if rf, ok := ret.Get(0).(func(*teachers.Domain) string); ok {
		r0 = rf(userDomain)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Update provides a mock function with given fields: id, blogDomain
func (_m *Usecase) Update(id string, blogDomain *teachers.Domain) teachers.Domain {
	ret := _m.Called(id, blogDomain)

	var r0 teachers.Domain
	if rf, ok := ret.Get(0).(func(string, *teachers.Domain) teachers.Domain); ok {
		r0 = rf(id, blogDomain)
	} else {
		r0 = ret.Get(0).(teachers.Domain)
	}

	return r0
}

type mockConstructorTestingTNewUsecase interface {
	mock.TestingT
	Cleanup(func())
}

// NewUsecase creates a new instance of Usecase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUsecase(t mockConstructorTestingTNewUsecase) *Usecase {
	mock := &Usecase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}