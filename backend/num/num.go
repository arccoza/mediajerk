package num

import (
	"cmp"
)

func Min[T cmp.Ordered](a, b T) T {
	if a < b {
		return a
	}

	return b
}

func MinOf[T cmp.Ordered](args ...T) T {
	a := args[0]
	for _, b := range args[1:] {
		if a < b {
			return a
		}

		a = b
	}

	return a
}

func Max[T cmp.Ordered](a, b T) T {
	if a > b {
		return a
	}

	return b
}

func MaxOf[T cmp.Ordered](args ...T) T {
	a := args[0]
	for _, b := range args[1:] {
		if a > b {
			return a
		}

		a = b
	}

	return a
}
