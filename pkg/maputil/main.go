package maputil

func Merge[M ~map[K]V, K comparable, V any](maps ...M) M {
	merged := make(M)

	for _, m := range maps {
		for k, v := range m {
			merged[k] = v
		}
	}

	return merged
}
