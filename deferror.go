package deferror

import (
	"errors"
)

// As Handles the error of a deferred function that returns an error. It will not overwrite the error to be
// returned by the function calling defer unless the current error is nil. If you want to capture both errors return
// both errors as separate variables. The calling function must initialize a named variable with the error interface
// in order for the As function to be able to return the error back to the parent of the caller function.
func As(f func() error, err *error) {
	if e := f(); e != nil && *err == nil {
		errors.As(e, err)
	}
}
