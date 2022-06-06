package libgo

func Map[T, U any](src []T, fn func(T) U) []U {
	dst := make([]U, len(src), len(src))
	for i, ele := range src {
		dst[i] = fn(ele)
	}
	return dst
}

func Filter[T any](src []T, fn func(T) bool) []T {
	dst := make([]T, 0, len(src))
	for _, ele := range src {
		if fn(ele) {
			dst = append(dst, ele)
		}
	}
	return dst
}
