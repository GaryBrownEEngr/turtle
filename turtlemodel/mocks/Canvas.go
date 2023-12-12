// Code generated by mockery v2.32.0. DO NOT EDIT.

package mocks

import (
	image "image"
	color "image/color"

	mock "github.com/stretchr/testify/mock"

	turtlemodel "github.com/GaryBrownEEngr/turtle/turtlemodel"
)

// Canvas is an autogenerated mock type for the Canvas type
type Canvas struct {
	mock.Mock
}

// ClearScreen provides a mock function with given fields: c
func (_m *Canvas) ClearScreen(c color.Color) {
	_m.Called(c)
}

// CreateNewSprite provides a mock function with given fields:
func (_m *Canvas) CreateNewSprite() turtlemodel.Sprite {
	ret := _m.Called()

	var r0 turtlemodel.Sprite
	if rf, ok := ret.Get(0).(func() turtlemodel.Sprite); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(turtlemodel.Sprite)
		}
	}

	return r0
}

// Exit provides a mock function with given fields:
func (_m *Canvas) Exit() {
	_m.Called()
}

// Fill provides a mock function with given fields: x, y, c
func (_m *Canvas) Fill(x int, y int, c color.Color) {
	_m.Called(x, y, c)
}

// GetHeight provides a mock function with given fields:
func (_m *Canvas) GetHeight() int {
	ret := _m.Called()

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// GetScreenshot provides a mock function with given fields:
func (_m *Canvas) GetScreenshot() image.Image {
	ret := _m.Called()

	var r0 image.Image
	if rf, ok := ret.Get(0).(func() image.Image); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(image.Image)
		}
	}

	return r0
}

// GetWidth provides a mock function with given fields:
func (_m *Canvas) GetWidth() int {
	ret := _m.Called()

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// PressedUserInput provides a mock function with given fields:
func (_m *Canvas) PressedUserInput() *turtlemodel.UserInput {
	ret := _m.Called()

	var r0 *turtlemodel.UserInput
	if rf, ok := ret.Get(0).(func() *turtlemodel.UserInput); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*turtlemodel.UserInput)
		}
	}

	return r0
}

// SetCartesianPixel provides a mock function with given fields: x, y, c
func (_m *Canvas) SetCartesianPixel(x int, y int, c color.Color) {
	_m.Called(x, y, c)
}

// SetPixel provides a mock function with given fields: x, y, c
func (_m *Canvas) SetPixel(x int, y int, c color.Color) {
	_m.Called(x, y, c)
}

// SubscribeToJustPressedUserInput provides a mock function with given fields:
func (_m *Canvas) SubscribeToJustPressedUserInput() chan *turtlemodel.UserInput {
	ret := _m.Called()

	var r0 chan *turtlemodel.UserInput
	if rf, ok := ret.Get(0).(func() chan *turtlemodel.UserInput); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(chan *turtlemodel.UserInput)
		}
	}

	return r0
}

// UnSubscribeToJustPressedUserInput provides a mock function with given fields: in
func (_m *Canvas) UnSubscribeToJustPressedUserInput(in chan *turtlemodel.UserInput) {
	_m.Called(in)
}

// NewCanvas creates a new instance of Canvas. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewCanvas(t interface {
	mock.TestingT
	Cleanup(func())
}) *Canvas {
	mock := &Canvas{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}