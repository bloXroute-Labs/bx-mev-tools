package ctxutil

import (
	"context"
	"io"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestValid(t *testing.T) {
	t.Run("cancel", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		require.NoError(t, Valid(ctx))

		cancel()
		require.ErrorIs(t, Valid(ctx), context.Canceled)
	})

	t.Run("cause", func(t *testing.T) {
		ctx, cancel := context.WithCancelCause(context.Background())
		defer cancel(nil)

		require.NoError(t, Valid(ctx))

		cancel(io.EOF)
		require.ErrorIs(t, Valid(ctx), context.Canceled)
		require.ErrorIs(t, context.Cause(ctx), io.EOF)
	})
}
