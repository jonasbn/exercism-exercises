package armstrong

import (
	"math"
	"strconv"
	"strings"
)

func IsNumber(n int) bool {

	s := strconv.Itoa(n)
	l := len(s)
	sum := 0

	c := strings.SplitAfter(s, "")

	for _, v := range c {
		t, err := strconv.Atoi(v)
		if err != nil {
			return false
		}
		sum += int(math.Pow(float64(t), float64(l)))
	}

	return sum == n
}
