package sieve

func Sieve(limit int) []int {

	var b []bool
	for i := 0; i <= limit; i++ {
		b = append(b, true)
	}

	for j := 2; j < len(b); j++ {
		if b[j] {
			for k := j; k < len(b); k += j {
				if k != j {
					b[k] = false
				}
			}
		}
	}

	var r []int
	for l := 2; l < len(b); l++ {
		if b[l] {
			r = append(r, l)
		}
	}

	return r
}
