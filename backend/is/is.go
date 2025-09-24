package is

import "cmp"

func Zero[T comparable](v T) bool {
	var z T
	return v == z
}

func Nil[T any](v T) bool {
	var a any = v
	return a == nil
}

func NaN[T cmp.Ordered](x T) bool {
	return x != x
}

// type Capable interface {
// 	~[]any | ~chan any
// }

// type Lenable interface {
// 	~[]any | ~map[any]any | ~chan any | ~string
// }

// type SliceLike interface {
// 	~[]any | ~chan any
// }

// type MapLike[K comparable] interface {
// 	~map[K]any
// }

// func Empty[T SliceLike | MapLike[K], K comparable](s T) bool {
// 	return len(s) == 0
// }

// func Full[T Capable](s T) bool {
// 	return len(s) == cap(s)
// }
