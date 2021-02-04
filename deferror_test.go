package deferror

import (
	"errors"
	"fmt"
	"testing"
)

var (
	logicFuncError = fmt.Errorf("logic function error")
	deferFuncError = fmt.Errorf("error coming from defer")
)

type TestDefer struct {
	name          string
	deferFunc     func() error
	logicFunc     func(func() error, error) (err error)
	passInError   error
	expectedError error
}

func Test(t *testing.T) {
	tests := []TestDefer{
		{
			name:          "defer returns error",
			deferFunc:     deferWithError,
			logicFunc:     logicFunction,
			passInError:   nil,
			expectedError: deferFuncError,
		},
		{
			name:          "no error",
			deferFunc:     deferWithoutError,
			logicFunc:     logicFunction,
			passInError:   nil,
			expectedError: nil,
		},
		{
			name:          "defer returns error and function returns error",
			deferFunc:     deferWithoutError,
			logicFunc:     logicFunction,
			passInError:   logicFuncError,
			expectedError: logicFuncError,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := test.logicFunc(test.deferFunc, test.passInError)
			if !errors.Is(test.expectedError, err) {
				t.Errorf("expected error does not match - expected: %s, actual: %s", test.expectedError.Error(), err.Error())
			}
		})
	}
}

func deferWithError() error {
	return deferFuncError
}

func deferWithoutError() error {
	return nil
}

func logicFunction(deferFunc func() error, passInError error) (err error) {
	defer As(deferFunc, &err)
	if passInError != nil {
		err = passInError
		return
	}
	return
}
