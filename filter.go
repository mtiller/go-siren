package siren

func Filter[T any](a []T, f func(x T) bool) []T {
	ret := []T{}
	for _, elem := range a {
		if f(elem) {
			ret = append(ret, elem)
		}
	}
	return ret
}
