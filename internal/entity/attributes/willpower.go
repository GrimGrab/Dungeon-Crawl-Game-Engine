package attributes

import (
	"SoB/internal/common"
	effects "SoB/internal/entity/attributes/effect"
)

const (
	MaxWillpower = 6
	MinWillpower = 1
)

type Willpower struct {
	willpower int

	effectManager *effects.EffectManager
}

func NewWillpower(willpower int, effectManager *effects.EffectManager) *Willpower {
	return &Willpower{
		willpower:     willpower,
		effectManager: effectManager,
	}
}

func (w *Willpower) BaseWillpower() int {
	return w.willpower
}

func (w *Willpower) Willpower() int {
	total := w.willpower + w.effectManager.CombatAttributeModifier(common.CombatAttributeWillpower)
	if total < MinWillpower {
		return MinWillpower
	} else if total > MaxWillpower {
		return MaxWillpower
	}
	return total
}
