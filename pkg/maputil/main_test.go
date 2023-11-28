package maputil

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMerge(t *testing.T) {
	m1 := map[int]int{1: 1, 2: 1, 3: 1}
	m2 := map[int]int{1: 2, 2: 2}
	m3 := map[int]int{1: 3}

	require.Equal(t, map[int]int{1: 3, 2: 2, 3: 1}, Merge(m1, m2, m3))
	require.Equal(t, map[int]int{1: 1, 2: 1, 3: 1}, Merge(m3, m2, m1))
}
