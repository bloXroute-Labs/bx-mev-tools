package errutil

import (
	"errors"
	"io"
)

// Close closes a closable and updates the error pointer if an error occurs
func Close(closable io.Closer, errPtr *error) {
	if closable == nil {
		return
	}

	if closeErr := closable.Close(); closeErr != nil {
		*errPtr = errors.Join(*errPtr, closeErr)
	}
}
