package errutil_test

import (
	"errors"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/bloXroute-Labs/bx-mev-tools/pkg/errutil"
)

var closeErr = errors.New("close error")

// MockCloser is a mock implementation of io.Closer
type MockCloser struct{ mock.Mock }

func (m *MockCloser) Close() error {
	args := m.Called()
	return args.Error(0)
}

func TestClose(t *testing.T) {
	tests := []struct {
		name        string
		closable    io.Closer
		closeError  error
		expectedErr *error
	}{
		{
			name:        "Successful Close",
			closable:    new(MockCloser),
			closeError:  nil,
			expectedErr: nil,
		},
		{
			name:        "Error On Close",
			closable:    new(MockCloser),
			closeError:  closeErr,
			expectedErr: &closeErr,
		},
		{
			name:        "Nil Closable",
			closable:    nil,
			closeError:  nil,
			expectedErr: nil,
		},
		{
			name:        "Nil Closable and Nil ErrPtr",
			closable:    nil,
			closeError:  nil,
			expectedErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var err error
			if tt.closable != nil && tt.name != "Nil Closable and Nil ErrPtr" {
				mockCloser := tt.closable.(*MockCloser)
				mockCloser.On("Close").Return(tt.closeError)
			}

			if tt.name != "Nil Closable and Nil ErrPtr" {
				errutil.Close(tt.closable, &err)
			} else {
				errutil.Close(tt.closable, nil) // Passing nil as the error pointer
			}

			if tt.expectedErr != nil {
				assert.EqualError(t, err, (*tt.expectedErr).Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
