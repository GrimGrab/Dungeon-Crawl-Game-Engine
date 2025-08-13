package dice

import (
	"crypto/rand"
	"math/big"
)

// Die represents a die with a set of sides, each side having an integer value.
type Die struct {
	sides []int
}

// Roll executes a roll of the die, returning a RollResult.
func (d Die) Roll() (RollResult, error) {
	n, err := rand.Int(rand.Reader, big.NewInt(int64(len(d.sides))))
	if err != nil {
		return RollResult{}, ErrRollFailed(d, err)
	}

	randomIndex := int(n.Int64())
	randomValue := d.sides[randomIndex]

	return RollResult{
		value: randomValue,
		die:   d,
	}, nil
}

// Sides returns the sides of the die.
func (d Die) Sides() []int {
	return d.sides
}

// NewStandardDie creates a new standard die with the specified number of sides from 1 to count.
// Example: NewStandardDie(6) creates a die with sides [1, 2, 3, 4, 5, 6].
func NewStandardDie(count int) (Die, error) {
	if count < 1 {
		return Die{}, ErrInvalidSides(count)
	}
	sides := make([]int, count)
	for i := 0; i < count; i++ {
		sides[i] = i + 1
	}
	return Die{sides: sides}, nil
}

// NewCustomDie creates a new die with custom sides.
// Example: NewCustomDie([]int{3, 3, 4, 4, 5, 6}) creates a die with those specific sides.
func NewCustomDie(sides []int) (Die, error) {
	if len(sides) < 1 {
		return Die{}, ErrInvalidSides(len(sides))
	}
	return Die{sides: sides}, nil
}
