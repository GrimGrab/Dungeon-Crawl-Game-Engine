package dice

type RollableSet interface {
	Roll() (RollResults, error)
}

type RollResult struct {
	value int
	die   Die
}

type RollResults []RollResult

func (r RollResults) Total() int {
	total := 0
	for _, result := range r {
		total += result.value
	}
	return total
}

func (r RollResults) Modify(mod func(results RollResults) RollResults) RollResults {
	return mod(r)
}

type Dice struct {
	dice []Die
}

func NewDice(dice ...Die) Dice {
	return Dice{dice: dice}
}

func (d Dice) Roll() (RollResults, error) {
	results := make(RollResults, len(d.dice))
	for i, die := range d.dice {
		result, err := die.Roll()
		if err != nil {
			return nil, ErrRollFailed(die, err)
		}
		results[i] = result
	}
	return results, nil
}
