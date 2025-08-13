package attributes

type Grit struct {
	grit    int
	maxGrit int
}

func NewGrit(initialGrit, maxGrit int) *Grit {
	return &Grit{
		grit:    initialGrit,
		maxGrit: maxGrit,
	}
}

func (g *Grit) Grit() int {
	return g.grit
}

func (g *Grit) MaxGrit() int {
	return g.maxGrit
}

func (g *Grit) IncreaseGrit(amount int) {
	if amount < 0 {
		return
	}
	g.grit += amount
	if g.grit > g.maxGrit {
		g.grit = g.maxGrit
	}
}

func (g *Grit) DecreaseGrit(amount int) {
	if amount < 0 {
		return
	}
	g.grit -= amount
	if g.grit < 0 {
		g.grit = 0
	}
}

func (g *Grit) IncreaseMaxGrit(amount int) {
	if amount < 0 {
		return
	}
	g.maxGrit += amount
}

func (g *Grit) DecreaseMaxGrit(amount int) {
	if amount < 0 {
		return
	}
	g.maxGrit -= amount
	if g.maxGrit < 0 {
		g.maxGrit = 0
	}
	if g.grit > g.maxGrit {
		g.grit = g.maxGrit
	}
}
