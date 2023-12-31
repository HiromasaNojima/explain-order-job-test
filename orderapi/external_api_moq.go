// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package orderapi

import (
	"sync"
)

// Ensure, that ExternalAPIMock does implement ExternalAPI.
// If this is not the case, regenerate this file with moq.
var _ ExternalAPI = &ExternalAPIMock{}

// ExternalAPIMock is a mock implementation of ExternalAPI.
//
// 	func TestSomethingThatUsesExternalAPI(t *testing.T) {
//
// 		// make and configure a mocked ExternalAPI
// 		mockedExternalAPI := &ExternalAPIMock{
// 			RequestOrderFunc: func(request OrderInfo) error {
// 				panic("mock out the RequestOrder method")
// 			},
// 		}
//
// 		// use mockedExternalAPI in code that requires ExternalAPI
// 		// and then make assertions.
//
// 	}
type ExternalAPIMock struct {
	// RequestOrderFunc mocks the RequestOrder method.
	RequestOrderFunc func(request OrderInfo) error

	// calls tracks calls to the methods.
	calls struct {
		// RequestOrder holds details about calls to the RequestOrder method.
		RequestOrder []struct {
			// Request is the request argument value.
			Request OrderInfo
		}
	}
	lockRequestOrder sync.RWMutex
}

// RequestOrder calls RequestOrderFunc.
func (mock *ExternalAPIMock) RequestOrder(request OrderInfo) error {
	if mock.RequestOrderFunc == nil {
		panic("ExternalAPIMock.RequestOrderFunc: method is nil but ExternalAPI.RequestOrder was just called")
	}
	callInfo := struct {
		Request OrderInfo
	}{
		Request: request,
	}
	mock.lockRequestOrder.Lock()
	mock.calls.RequestOrder = append(mock.calls.RequestOrder, callInfo)
	mock.lockRequestOrder.Unlock()
	return mock.RequestOrderFunc(request)
}

// RequestOrderCalls gets all the calls that were made to RequestOrder.
// Check the length with:
//     len(mockedExternalAPI.RequestOrderCalls())
func (mock *ExternalAPIMock) RequestOrderCalls() []struct {
	Request OrderInfo
} {
	var calls []struct {
		Request OrderInfo
	}
	mock.lockRequestOrder.RLock()
	calls = mock.calls.RequestOrder
	mock.lockRequestOrder.RUnlock()
	return calls
}
