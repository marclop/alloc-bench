package alloctest

type T struct {
	A, B, C string
	I       int
}

func pointer(t *T) string {
	return t.A
}

func plain(t T) string {
	return t.A
}
