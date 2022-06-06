package libgo

func Must(err error) {
	if err != nil {
		panic(err)
	}
}

func MustOne[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}
	return v
}

func MustTwo[T, U any](v1 T, v2 U, err error) (T, U) {
	if err != nil {
		panic(err)
	}
	return v1, v2
}
