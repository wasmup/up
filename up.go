package up

// Grouping inserts the "_" separator between every 3 chars e.g.:
// "-1234567" returns "-1_234_567"
func Grouping(st string) string {
	if st == "" {
		return st
	}
	sign, s := "", ""
	if st[0] == '-' {
		sign, st = "-", st[1:]
	}
	rs := []rune(st)
	for i := range rs {
		if ((i % 3) == 0) && (i != 0) {
			s = "_" + s
		}
		s = string(rs[len(rs)-1-i]) + s
	}
	return sign + s
}

func Abs[T Number](v T) T {
	if v < 0 {
		v = -v
	}
	return v
}

func Min[T Number](a, b T) T {
	if a < b {
		return a
	}
	return b
}

func Max[T Number](a, b T) T {
	if a > b {
		return a
	}
	return b
}

type Number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		// byte | // alias for uint8
		// rune | // alias for int32 // represents a Unicode code point
		~float32 | ~float64
	//  ~complex64 | ~complex128
}
