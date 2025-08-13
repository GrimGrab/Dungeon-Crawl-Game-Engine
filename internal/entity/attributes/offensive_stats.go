package attributes

import (
	"SoB/internal/common"
	effects "SoB/internal/entity/attributes/effect"
)

type OffensiveStats struct {
	rangedToHit int
	meleeToHit  int
	combat      int

	effectManager *effects.EffectManager
}

func NewOffensiveStats(rangedToHit, meleeToHit, combat int, effectManager *effects.EffectManager) *OffensiveStats {
	return &OffensiveStats{
		rangedToHit: rangedToHit,
		meleeToHit:  meleeToHit,
		combat:      combat,

		effectManager: effectManager,
	}
}

func (os *OffensiveStats) BaseRangedToHit() int {
	return os.rangedToHit
}

func (os *OffensiveStats) BaseMeleeToHit() int {
	return os.meleeToHit
}

func (os *OffensiveStats) BaseCombat() int {
	return os.combat
}

func (os *OffensiveStats) RangedToHit() int {
	modifier := os.effectManager.CombatAttributeModifier(common.CombatAttributeRangedToHit)
	if modifier > 0 {
		return modifier
	}
	return os.rangedToHit
}

func (os *OffensiveStats) MeleeToHit() int {
	modifier := os.effectManager.CombatAttributeModifier(common.CombatAttributeMeleeToHit)
	if modifier > 0 {
		return modifier
	}
	return os.meleeToHit
}

func (os *OffensiveStats) Combat() int {
	return os.combat + os.effectManager.CombatAttributeModifier(common.CombatAttributeCombat)
}
