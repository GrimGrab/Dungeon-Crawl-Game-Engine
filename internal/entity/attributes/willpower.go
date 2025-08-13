package attributes

const (
	MaxWillpower = 6
	MinWillpower = 1
)

type Willpower struct {
	willpower int
}

func NewWillpower(willpower int) *Willpower {
	return &Willpower{
		willpower: willpower,
	}
}

func (w *Willpower) BaseWillpower() int {
	return w.willpower
}
