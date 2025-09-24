package non

func Zero[T comparable](a, b T) T {
	var z T
	if a != z {
		return a
	}
	return b
}

func Zeros[T comparable](vals ...T) T {
	var z T
	for _, v := range vals {
		if v != z {
			return v
		}
	}
	return z
}

func Nil[T any](a, b T) T {
	var _a any = a
	if _a != nil {
		return _a.(T)
	}

	return b
}

func Nils[T any](vals ...T) T {
	var n T
	for _, v := range vals {
		var a any = v
		if a != nil {
			return a.(T)
		}
	}

	return n
}

func Empty[T []any](a, b T) T {
	if len(a) > 0 {
		return a
	}

	return b
}

func Empties[T []any](vals ...T) T {
	for _, v := range vals {
		if len(v) > 0 {
			return v
		}
	}

	return nil
}
