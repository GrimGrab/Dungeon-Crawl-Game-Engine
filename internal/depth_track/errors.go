package depth_track

import "fmt"

type DepthError struct {
	Op      string
	Message string
	Err     error
}

func (e *DepthError) Error() string {
	return fmt.Sprintf("Operation: %s; Message: %s", e.Op, e.Message)
}

func (e *DepthError) Unwrap() error {
	return e.Err
}

func ErrLevelOutOfBounds(level, minLevel, maxLevel int) *DepthError {
	return &DepthError{
		Op:      "set level",
		Message: fmt.Sprintf("level %d is out of bounds (%d - %d)", level, minLevel, maxLevel),
	}
}

func ErrMinGreaterThanMax(minLevel, maxLevel int) *DepthError {
	return &DepthError{
		Op:      "create",
		Message: fmt.Sprintf("minLevel %d cannot be greater than maxLevel %d", minLevel, maxLevel),
	}
}
