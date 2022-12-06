package mocks

import (
	"context"

	"github.com/plus100kt/goserver/gag/model"
	"github.com/stretchr/testify/mock"
)

// MockDeviceService is a mock type for model.DeviceService
type MockDeviceService struct {
	mock.Mock
}

func (m *MockDeviceService) Register(ctx context.Context, uuid string) (*model.Device, error) {
	// args that will be passed to "Return" in the tests, when function
	// is called with a uid. Hence the name "ret"
	ret := m.Called(ctx, uuid)

	// first value passed to "Return"
	var r0 *model.Device
	if ret.Get(0) != nil {
		// we can just return this if we know we won't be passing function to "Return"
		r0 = ret.Get(0).(*model.Device)
	}

	var r1 error

	if ret.Get(1) != nil {
		r1 = ret.Get(1).(error)
	}

	return r0, r1
}

func (m *MockDeviceService) Get(ctx context.Context, uuid string) (*model.Device, error) {
	// args that will be passed to "Return" in the tests, when function
	// is called with a uid. Hence the name "ret"
	ret := m.Called(ctx, uuid)

	// first value passed to "Return"
	var r0 *model.Device
	if ret.Get(0) != nil {
		// we can just return this if we know we won't be passing function to "Return"
		r0 = ret.Get(0).(*model.Device)
	}

	var r1 error

	if ret.Get(1) != nil {
		r1 = ret.Get(1).(error)
	}

	return r0, r1
}
