package ptrutil

// Ptr returns pointer to provided value.
func Ptr[T any](val T) *T { return &val }
