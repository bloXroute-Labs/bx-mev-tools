package ctxutil_test

import (
	"context"
	"io"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/bloXroute-Labs/bx-mev-tools/pkg/ctxutil"
)

func TestValid(t *testing.T) {
	rootCtx, rootCancel := context.WithCancel(context.Background())
	defer rootCancel()

	tests := []struct {
		name        string
		contextFunc func() context.Context
		expectedErr error
		cause       error
	}{
		{
			name:        "Valid_Context",
			contextFunc: func() context.Context { return rootCtx },
		},
		{
			name: "Canceled_Context",
			contextFunc: func() context.Context {
				ctx, cancel := context.WithCancel(rootCtx)
				cancel()
				return ctx
			},
			expectedErr: context.Canceled,
		},
		{
			name: "Deadline_Exceeded_Context",
			contextFunc: func() context.Context {
				ctx, cancel := context.WithTimeout(rootCtx, time.Nanosecond)
				defer cancel()
				time.Sleep(time.Millisecond) // Ensure the deadline is exceeded
				return ctx
			},
			expectedErr: context.DeadlineExceeded,
		},
		{
			name: "Context_With_Cancel_Cause_Nil",
			contextFunc: func() context.Context {
				ctx, cancel := context.WithCancelCause(rootCtx)
				cancel(nil)
				return ctx
			},
			expectedErr: context.Canceled,
		},
		{
			name: "Context_With_Cancel_Cause_EOF",
			contextFunc: func() context.Context {
				ctx, cancel := context.WithCancelCause(rootCtx)
				cancel(io.EOF)
				return ctx
			},
			expectedErr: context.Canceled,
			cause:       io.EOF,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := tt.contextFunc()
			err := ctxutil.Valid(ctx)

			if tt.expectedErr != nil {
				assert.ErrorIs(t, err, tt.expectedErr)
			} else {
				assert.NoError(t, err)
			}

			if tt.cause != nil {
				assert.ErrorIs(t, context.Cause(ctx), tt.cause)
			}
		})
	}
}
