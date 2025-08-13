package dice

import "fmt"

var (
	D3    = Must(NewStandardDie(3))
	D4    = Must(NewStandardDie(4))
	D6    = Must(NewStandardDie(6))
	D8    = Must(NewStandardDie(8))
	D10   = Must(NewStandardDie(10))
	D12   = Must(NewStandardDie(12))
	D20   = Must(NewStandardDie(20))
	Peril = Must(NewCustomDie([]int{3, 3, 4, 4, 5, 6}))
)

// Must is a helper function that panics if err is not nil, otherwise returns the value
func Must(die Die, err error) Die {
	if err != nil {
		panic(fmt.Sprintf("dice creation failed: %v", err))
	}
	return die
}
