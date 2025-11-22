package grains

import "errors"

func Square(number int) (uint64, error) {
	var result uint64

	if number < 1 || number > 64 {
		return uint64(0), errors.New("Please specify a positive number between 1 and 64")
	}

	if number == 1 {
		return uint64(number), nil
	}

	result, _ = Square(number - 1)

	return result * 2, nil
}

func Total() uint64 {
	var total uint64 = 0
	i := 1
	for i <= 64 {
		t, _ := Square(i)
		total += t
		i++
	}
	return uint64(total)
}
