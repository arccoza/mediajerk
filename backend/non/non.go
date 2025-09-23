package non

func Zero[T comparable](vals ...T) T {
	var z T
	for _, v := range vals {
		if v != z {
			return v
		}
	}
	return z
}

func Nil[T any](vals ...T) T {
	var n T
	for _, v := range vals {
		var a any = v
		if a != nil {
			return a.(T)
		}
	}

	return n
}

func Empty[T []any](vals ...T) T {
	for _, v := range vals {
		if len(v) > 0 {
			return v
		}
	}

	return nil
}
