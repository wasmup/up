package up

func Grouping(st string) (s string) {
	if st[0] == '-' {
		return "-" + Grouping(st[1:])
	}
	rs := []rune(st)
	for i := range rs {
		if ((i % 3) == 0) && (i != 0) {
			s = "_" + s
		}
		s = string(rs[len(rs)-1-i]) + s
	}
	return
}
