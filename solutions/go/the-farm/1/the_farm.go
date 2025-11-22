package thefarm

import (
	"errors"
	"fmt"
)

// See types.go for the types defined for this exercise.

//  SillyNephewError type for handling negative numbers of cows.
type SillyNephewError struct {
	Cows int
}

func (e *SillyNephewError) Error() string {
	return fmt.Sprintf("silly nephew, there cannot be %d cows", e.Cows)
}

// DivideFood computes the fodder amount per cow for the given cows.
func DivideFood(weightFodder WeightFodder, cows int) (float64, error) {
	fodder, error := weightFodder.FodderAmount()

	if error == ErrScaleMalfunction {
		fodder *= 2
		error = nil
	}

	if error != nil {
		return 0, error
	}

	if fodder < 0 {
		return 0, errors.New("negative fodder")
	}

	if cows == 0 {
		return 0.0, errors.New("division by zero")
	}

	if cows < 0 {
		return 0.0, &SillyNephewError{cows}
	}

	return fodder / float64(cows), nil
}
