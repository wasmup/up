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
