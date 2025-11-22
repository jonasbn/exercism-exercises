package collatzconjecture

import "errors"

func CollatzConjecture(number int) (int, error) {
	if number < 1 {
		return 0, errors.New("zero or negative numbers not valid input")
	}

	steps := 0

	for number != 1 {
		if number%2 == 0 {
			number = number / 2
		} else {
			number = number*3 + 1
		}
		steps++
	}

	return steps, nil
}
