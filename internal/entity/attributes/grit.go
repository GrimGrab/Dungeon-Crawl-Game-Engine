package attributes

import (
	"SoB/internal/common"
	effects "SoB/internal/entity/attributes/effect"
)

type Grit struct {
	grit    int
	maxGrit int

	effectManager *effects.EffectManager
}

func NewGrit(initialGrit, maxGrit int, effectsManager *effects.EffectManager) *Grit {
	return &Grit{
		grit:    initialGrit,
		maxGrit: maxGrit,

		effectManager: effectsManager,
	}
}

func (g *Grit) Grit() int {
	return g.grit
}

func (g *Grit) BaseMaxGrit() int {
	return g.maxGrit
}

func (g *Grit) MaxGrit() int {
	return g.maxGrit + g.effectManager.AttributeModifier(common.AttributeGrit)
}

func (g *Grit) IncreaseGrit(amount int) {
	if amount < 0 {
		return
	}
	g.grit += amount
	if g.grit > g.maxGrit+g.effectManager.AttributeModifier(common.AttributeGrit) {
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
