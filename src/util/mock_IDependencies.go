// Code generated by mockery v1.0.0. DO NOT EDIT.

package util

import (
	ghw "github.com/jaypipes/ghw"
	mock "github.com/stretchr/testify/mock"

	netlink "github.com/vishvananda/netlink"

	os "os"
)

// MockIDependencies is an autogenerated mock type for the IDependencies type
type MockIDependencies struct {
	mock.Mock
}

// Abs provides a mock function with given fields: path
func (_m *MockIDependencies) Abs(path string) (string, error) {
	ret := _m.Called(path)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(path)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(path)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Block provides a mock function with given fields: opts
func (_m *MockIDependencies) Block(opts ...*ghw.WithOption) (*ghw.BlockInfo, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ghw.BlockInfo
	if rf, ok := ret.Get(0).(func(...*ghw.WithOption) *ghw.BlockInfo); ok {
		r0 = rf(opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ghw.BlockInfo)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(...*ghw.WithOption) error); ok {
		r1 = rf(opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// EvalSymlinks provides a mock function with given fields: path
func (_m *MockIDependencies) EvalSymlinks(path string) (string, error) {
	ret := _m.Called(path)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(path)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(path)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Execute provides a mock function with given fields: command, args
func (_m *MockIDependencies) Execute(command string, args ...string) (string, string, int) {
	_va := make([]interface{}, len(args))
	for _i := range args {
		_va[_i] = args[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, command)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 string
	if rf, ok := ret.Get(0).(func(string, ...string) string); ok {
		r0 = rf(command, args...)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 string
	if rf, ok := ret.Get(1).(func(string, ...string) string); ok {
		r1 = rf(command, args...)
	} else {
		r1 = ret.Get(1).(string)
	}

	var r2 int
	if rf, ok := ret.Get(2).(func(string, ...string) int); ok {
		r2 = rf(command, args...)
	} else {
		r2 = ret.Get(2).(int)
	}

	return r0, r1, r2
}

// Hostname provides a mock function with given fields:
func (_m *MockIDependencies) Hostname() (string, error) {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Interfaces provides a mock function with given fields:
func (_m *MockIDependencies) Interfaces() ([]Interface, error) {
	ret := _m.Called()

	var r0 []Interface
	if rf, ok := ret.Get(0).(func() []Interface); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]Interface)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LinkByName provides a mock function with given fields: name
func (_m *MockIDependencies) LinkByName(name string) (netlink.Link, error) {
	ret := _m.Called(name)

	var r0 netlink.Link
	if rf, ok := ret.Get(0).(func(string) netlink.Link); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(netlink.Link)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Product provides a mock function with given fields: opts
func (_m *MockIDependencies) Product(opts ...*ghw.WithOption) (*ghw.ProductInfo, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ghw.ProductInfo
	if rf, ok := ret.Get(0).(func(...*ghw.WithOption) *ghw.ProductInfo); ok {
		r0 = rf(opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ghw.ProductInfo)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(...*ghw.WithOption) error); ok {
		r1 = rf(opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReadDir provides a mock function with given fields: dirname
func (_m *MockIDependencies) ReadDir(dirname string) ([]os.FileInfo, error) {
	ret := _m.Called(dirname)

	var r0 []os.FileInfo
	if rf, ok := ret.Get(0).(func(string) []os.FileInfo); ok {
		r0 = rf(dirname)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]os.FileInfo)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(dirname)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReadFile provides a mock function with given fields: fname
func (_m *MockIDependencies) ReadFile(fname string) ([]byte, error) {
	ret := _m.Called(fname)

	var r0 []byte
	if rf, ok := ret.Get(0).(func(string) []byte); ok {
		r0 = rf(fname)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(fname)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RouteList provides a mock function with given fields: link, family
func (_m *MockIDependencies) RouteList(link netlink.Link, family int) ([]netlink.Route, error) {
	ret := _m.Called(link, family)

	var r0 []netlink.Route
	if rf, ok := ret.Get(0).(func(netlink.Link, int) []netlink.Route); ok {
		r0 = rf(link, family)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]netlink.Route)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(netlink.Link, int) error); ok {
		r1 = rf(link, family)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Stat provides a mock function with given fields: fname
func (_m *MockIDependencies) Stat(fname string) (os.FileInfo, error) {
	ret := _m.Called(fname)

	var r0 os.FileInfo
	if rf, ok := ret.Get(0).(func(string) os.FileInfo); ok {
		r0 = rf(fname)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(os.FileInfo)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(fname)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
