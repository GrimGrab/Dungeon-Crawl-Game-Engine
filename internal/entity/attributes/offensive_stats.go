package attributes

import "SoB/internal/common"

type OffensiveStats struct {
	rangedToHit int
	meleeToHit  int
	combat      int
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

type OffensiveStatsEffects struct {
	name        string
	description string

	rangedToHitChange int
	meleeToHitChange  int
	combatChange      int

	durationType common.DurationType
	duration     int
}
