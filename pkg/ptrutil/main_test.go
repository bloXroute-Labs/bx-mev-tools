package ptrutil_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/bloXroute-Labs/bx-mev-tools/pkg/ptrutil"
)

type CustomType struct{ Field string }

func TestPtr(t *testing.T) {
	testCases := []struct {
		name  string
		value any
	}{
		{"Int", 42},
		{"String", "hello"},
		{"Float", 3.14},
		{"Bool", true},
		{"Byte", byte('a')},
		{"Rune", rune('🚀')},
		{"Slice", []int{1, 2, 3}},
		{"Map", map[string]int{"one": 1, "two": 2}},
		{"CustomType", CustomType{Field: "Test"}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := ptrutil.Ptr(tc.value)
			assert.NotNil(t, got)
			assert.Equal(t, &tc.value, got)
		})
	}
}
