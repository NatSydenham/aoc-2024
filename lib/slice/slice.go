package slice

func RemoveAt[T any](slice []T, s int) []T {
	result := make([]T, len(slice))
	copy(result, slice)
	return append(result[:s], result[s+1:]...)
}

func Map[T, V any](ts []T, fn func(T) V) []V {
	result := make([]V, len(ts))
	for i, t := range ts {
		result[i] = fn(t)
	}
	return result
}

func IndexOf[T comparable](el T, ts []T) int {
	for k, v := range ts {
		if el == v {
			return k
		}
	}
	return -1 //not found.
}
