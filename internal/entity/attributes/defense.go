package attributes

import (
	"SoB/internal/common"
	effects "SoB/internal/entity/attributes/effect"
)

const MaxDefense = 6
const MinDefense = 1

type Defense struct {
	defense       int
	effectManager *effects.EffectManager
}

func NewDefense(defense int) *Defense {
	return &Defense{
		defense: defense,
	}
}

func (d *Defense) BaseDefense() int {
	return d.defense
}

func (d *Defense) Defense() int {
	total := d.defense + d.effectManager.CombatAttributeModifier(common.CombatAttributeDefense)
	if total < MinDefense {
		return MinDefense
	} else if total > MaxDefense {
		return MaxDefense
	}
	return total
}
