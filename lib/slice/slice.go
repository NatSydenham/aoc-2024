package slice

func RemoveAt[T any](slice []T, s int) []T {
	result := make([]T, len(slice))
	copy(result, slice)
	return append(result[:s], result[s+1:]...)
}