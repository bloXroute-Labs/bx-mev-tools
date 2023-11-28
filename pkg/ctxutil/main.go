package ctxutil

import "context"

// Valid checks if the context is valid and returns an error if not
func Valid(ctx context.Context) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		return nil
	}
}
