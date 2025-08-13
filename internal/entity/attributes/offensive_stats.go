package attributes

import "SoB/internal/common"

type OffensiveStats struct {
	rangedToHit int
	meleeToHit  int
	combat      int

	effects []*OffensiveStatsEffects
}

func NewOffensiveStats(rangedToHit, meleeToHit, combat int) *OffensiveStats {
	return &OffensiveStats{
		rangedToHit: rangedToHit,
		meleeToHit:  meleeToHit,
		combat:      combat,
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

func (os *OffensiveStats) TotalRangedToHit() int {
	total := os.rangedToHit
	for _, effect := range os.effects {
		total += effect.RangedToHitChange()
	}
	return total
}

func (os *OffensiveStats) TotalMeleeToHit() int {
	total := os.meleeToHit
	for _, effect := range os.effects {
		total += effect.MeleeToHitChange()
	}
	return total
}

func (os *OffensiveStats) TotalCombat() int {
	total := os.combat
	for _, effect := range os.effects {
		total += effect.CombatChange()
	}
	return total
}

func (os *OffensiveStats) Effects() []*OffensiveStatsEffects {
	return os.effects
}

func (os *OffensiveStats) AddEffect(effect *OffensiveStatsEffects) {
	os.effects = append(os.effects, effect)
}

func (os *OffensiveStats) RemoveEffect(effect *OffensiveStatsEffects) {
	for i, e := range os.effects {
		if e == effect {
			os.effects = append(os.effects[:i], os.effects[i+1:]...)
			break
		}
	}
}

func (os *OffensiveStats) DecrementEffectDurations(durationType common.DurationType) {
	for i := len(os.effects) - 1; i >= 0; i-- {
		effect := os.effects[i]
		if effect.DurationType() == durationType {
			effect.DecreaseDuration()
			if effect.Expired() {
				os.effects[i] = os.effects[len(os.effects)-1]
				os.effects = os.effects[:len(os.effects)-1]
			}
		}
	}
}

type OffensiveStatsEffects struct {
	name        string
	description string

	rangedToHitChange int
	meleeToHitChange  int
	combatChange      int

	durationType common.DurationType
	duration     int
}

func NewOffensiveStatsEffects(name, description string, rangedToHit, meleeToHit, combat int, durationType common.DurationType, duration int) *OffensiveStatsEffects {
	return &OffensiveStatsEffects{
		name:        name,
		description: description,

		rangedToHitChange: rangedToHit,
		meleeToHitChange:  meleeToHit,
		combatChange:      combat,

		durationType: durationType,
		duration:     duration,
	}
}

func (ose *OffensiveStatsEffects) Name() string {
	return ose.name
}

func (ose *OffensiveStatsEffects) Description() string {
	return ose.description
}

func (ose *OffensiveStatsEffects) RangedToHitChange() int {
	return ose.rangedToHitChange
}

func (ose *OffensiveStatsEffects) MeleeToHitChange() int {
	return ose.meleeToHitChange
}

func (ose *OffensiveStatsEffects) CombatChange() int {
	return ose.combatChange
}

func (ose *OffensiveStatsEffects) DurationType() common.DurationType {
	return ose.durationType
}

func (ose *OffensiveStatsEffects) Duration() int {
	return ose.duration
}

func (ose *OffensiveStatsEffects) DecreaseDuration() {
	if ose.duration > 0 {
		ose.duration--
	}
}

func (ose *OffensiveStatsEffects) Expired() bool {
	return ose.duration == 0
}
