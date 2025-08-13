package entity

type CharacterError struct {
	Message string
	Err     error
}

func (e *CharacterError) Error() string {
	return e.Message
}

func (e *CharacterError) Unwrap() error {
	return e.Err
}

func ErrInvalidEntityConstruction(message string) *CharacterError {
	return &CharacterError{
		Message: message,
	}
}
