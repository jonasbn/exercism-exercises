package lsproduct

import (
	"fmt"
	"strconv"
	"strings"
)

func LargestSeriesProduct(digits string, span int) (int64, error) {

	if span < 0 {
		err := fmt.Errorf("span must not be negative")
		return 0, err
	}

	if span > len(digits) {
		err := fmt.Errorf("span must be smaller than string length")
		return 0, err
	}

	if len(digits) < 1 {
		return 1, nil
	}

	var t string
	var highestProduct int64

	for i := 0; i < len(digits); i++ {
		if i+span > len(digits) {
			break
		}
		t = digits[i : i+span]

		p, err := product(t)
		if err != nil {
			return 0, err
		}

		if p > highestProduct {
			highestProduct = p
		}
	}

	return int64(highestProduct), nil
}

func product(ss string) (int64, error) {

	s := strings.SplitAfter(ss, "")

	var product int64 = 1

	for _, k := range s {
		o, err := strconv.Atoi(k)
		if err != nil {
			return 0, err
		}
		product = product * int64(o)
	}

	return product, nil
}
