package utils

func Map[T interface{}, R interface{}](arr []T, f func(T) R) []R {
	rs := make([]R, len(arr))
	for i, v := range arr {
		rs[i] = f(v)
	}
	return rs
}

func Filter[T interface{}](source []T, test func(T) bool) (dest []T) {
	for _, el := range source {
		if test(el) {
			dest = append(dest, el)
		}
	}
	return
}

func FindFirst[T any](source []T, test func(T) bool) (result T, found bool) {
	dest := Filter(source, test)
	if dest == nil || len(dest) == 0 {
		var rs T
		return rs, false
	} else {
		return dest[0], true
	}
}

func Count[T interface{}](source []T, test func(T) bool) int {
	return len(Filter(source, test))
}