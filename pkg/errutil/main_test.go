package errutil

import (
	"errors"
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

var errMockClose = errors.New("mock close error")

// mockClosable is used to test the Close function
type mockClosable struct{}

// Close overrides the Close method to return an error
func (m *mockClosable) Close() error { return errMockClose }

func TestClose(t *testing.T) {
	t.Run("closable nil", func(t *testing.T) {
		var err = io.EOF

		Close(nil, &err)
		require.Error(t, err)
		require.ErrorIs(t, err, io.EOF)

		err = io.ErrUnexpectedEOF
		Close(nil, &err)
		require.Error(t, err)
		require.NotErrorIs(t, err, io.EOF)
		require.ErrorIs(t, err, io.ErrUnexpectedEOF)
	})

	t.Run("closable not nil", func(t *testing.T) {
		file, err := os.CreateTemp("", "closable")
		require.NoError(t, err)
		require.Implements(t, (*io.Closer)(nil), file)

		defer func() { require.NoError(t, os.Remove(file.Name())) }()

		Close(file, &err)
		require.NoError(t, err)

		closable := new(mockClosable)
		require.Implements(t, (*io.Closer)(nil), closable)

		Close(closable, &err)
		require.ErrorIs(t, err, errMockClose)
		require.NotErrorIs(t, err, io.EOF)

		err = io.EOF
		require.ErrorIs(t, err, io.EOF)
		require.NotErrorIs(t, err, errMockClose)

		Close(closable, &err)
		require.ErrorIs(t, err, io.EOF)
		require.ErrorIs(t, err, errMockClose)
	})

	t.Run("err nil", func(t *testing.T) {
		var err error

		require.NoError(t, err)

		Close(new(mockClosable), &err)
		require.ErrorIs(t, err, errMockClose)
	})

	t.Run("closable nil and err nil", func(t *testing.T) {
		var err error

		require.NoError(t, err)

		Close(nil, &err)
		require.NoError(t, err)
	})
}
