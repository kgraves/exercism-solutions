package thefarm

import "errors"
import "fmt"

func DivideFood(fc FodderCalculator, numCows int) (float64, error) {
    amount, err := fc.FodderAmount(numCows)
    if err != nil {
        return 0.0, err
    }

    factor, err := fc.FatteningFactor()
    if err != nil {
        return 0.0, err
    }

    return (amount * factor) / float64(numCows), nil
}

func ValidateInputAndDivideFood(fc FodderCalculator, numCows int) (float64, error) {
    if numCows <= 0 {
        return 0.0, errors.New("invalid number of cows")
    }

    return DivideFood(fc, numCows)
}

type InvalidCowsError struct {
    numCows int
    message string
}

func (e *InvalidCowsError) Error() string {
    return fmt.Sprintf("%d cows are invalid: %s", e.numCows, e.message)
}

func ValidateNumberOfCows(numCows int) error {
    if numCows < 0 {
        return &InvalidCowsError{numCows, "there are no negative cows"}
    } else if numCows == 0 {
        return &InvalidCowsError{numCows, "no cows don't need food"}
    } else {
        return nil
    }
}
