package dice

import "fmt"

type RollError struct {
	Op      string
	Die     Die
	Message string
	Err     error
}

func (e *RollError) Error() string {
	return fmt.Sprintf("Operation: %s;Sides %d; Message %s", e.Op, e.Die.Sides(), e.Message)
}

func (e *RollError) Unwrap() error {
	return e.Err
}

func ErrInvalidSides(sides int) *RollError {
	return &RollError{
		Op:      "create",
		Message: fmt.Sprintf("invalid number of sides: %d (must be >= 1)", sides),
	}
}

func ErrRollFailed(die Die, err error) *RollError {
	return &RollError{
		Op:      "roll",
		Die:     die,
		Message: fmt.Sprintf("failed to roll die, %d", die.Sides()),
		Err:     err,
	}
}
