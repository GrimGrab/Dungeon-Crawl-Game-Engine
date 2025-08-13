package entity

import (
	"SoB/internal/entity/attributes"
	"fmt"
)

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

func ErrInvalidName() *CharacterError {
	return &CharacterError{
		Message: "invalid name: name cannot be empty",
	}
}

func ErrInvalidHealth() *CharacterError {
	return &CharacterError{
		Message: fmt.Sprintf("invalid health: health must be greater than %d", minimumHealth),
	}
}

func ErrInvalidSanity() *CharacterError {
	return &CharacterError{
		Message: fmt.Sprintf("invalid sanity: sanity must be at least %d", minimumSanity),
	}
}

func ErrInvalidGrit() *CharacterError {
	return &CharacterError{
		Message: fmt.Sprintf("invalid grit: max grit must be atleast %d", minimumGrit),
	}
}

func ErrInvalidToHitRoll() *CharacterError {
	return &CharacterError{
		Message: fmt.Sprintf("invalid to-hit roll: to-hit roll must be between %d and %d", minimumToHitRoll, maximumToHitRoll),
	}
}

func ErrInvalidCombat() *CharacterError {
	return &CharacterError{
		Message: fmt.Sprintf("invalid combat: combat must be at least %d", minimumCombat),
	}
}

func ErrInvalidStats() *CharacterError {
	return &CharacterError{
		Message: fmt.Sprintf("invalid stats: all stats must be at least %d", attributes.MinimumStatValue),
	}
}
