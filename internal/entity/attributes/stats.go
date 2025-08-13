package attributes

import "SoB/internal/common"

const MinimumStatValue = 1

// Stats represents the core attributes of a character, including base values and any active effects.
type Stats struct {
	agility  int
	cunning  int
	spirit   int
	strength int
	lore     int
	luck     int

	effects []*StatEffect
}

// NewStats creates a new Stats instance with the given base attribute values.
func NewStats(agility, cunning, spirit, strength, lore, luck int) *Stats {
	return &Stats{
		agility:  agility,
		cunning:  cunning,
		spirit:   spirit,
		strength: strength,
		lore:     lore,
		luck:     luck,
		effects:  make([]*StatEffect, 0),
	}
}

// BaseAgility returns the base agility value without any effects.
func (s *Stats) BaseAgility() int {
	return s.agility
}

// BaseCunning returns the base cunning value without any effects.
func (s *Stats) BaseCunning() int {
	return s.cunning
}

// BaseSpirit returns the base spirit value without any effects.
func (s *Stats) BaseSpirit() int {
	return s.spirit
}

// BaseStrength returns the base strength value without any effects.
func (s *Stats) BaseStrength() int {
	return s.strength
}

// BaseLore returns the base lore value without any effects.
func (s *Stats) BaseLore() int {
	return s.lore
}

// BaseLuck returns the base luck value without any effects.
func (s *Stats) BaseLuck() int {
	return s.luck
}

// calculateTotal computes the total value of a stat by adding base value and all effects.
func (s *Stats) calculateTotal(baseStat int, getEffectStat func(*StatEffect) int) int {
	total := baseStat
	for _, effect := range s.effects {
		total += getEffectStat(effect)
	}
	if total < MinimumStatValue {
		return MinimumStatValue
	}
	return total
}

// TotalAgility returns the total agility value including all effects.
func (s *Stats) TotalAgility() int {
	return s.calculateTotal(s.agility, func(e *StatEffect) int { return e.agility })
}

// TotalCunning returns the total cunning value including all effects.
func (s *Stats) TotalCunning() int {
	return s.calculateTotal(s.cunning, func(e *StatEffect) int { return e.cunning })
}

// TotalSpirit returns the total spirit value including all effects.
func (s *Stats) TotalSpirit() int {
	return s.calculateTotal(s.spirit, func(e *StatEffect) int { return e.spirit })
}

// TotalStrength returns the total strength value including all effects.
func (s *Stats) TotalStrength() int {
	return s.calculateTotal(s.strength, func(e *StatEffect) int { return e.strength })
}

// TotalLore returns the total lore value including all effects.
func (s *Stats) TotalLore() int {
	return s.calculateTotal(s.lore, func(e *StatEffect) int { return e.lore })
}

// TotalLuck returns the total luck value including all effects.
func (s *Stats) TotalLuck() int {
	return s.calculateTotal(s.luck, func(e *StatEffect) int { return e.luck })
}

// Effects returns a slice of all active StatEffect on the Stats.
func (s *Stats) Effects() []*StatEffect {
	return s.effects
}

// AddEffect adds a new StatEffect to the Stats.
func (s *Stats) AddEffect(effect *StatEffect) {
	s.effects = append(s.effects, effect)
}

// RemoveEffect removes a specific StatEffect from the Stats.
func (s *Stats) RemoveEffect(effect *StatEffect) {
	for i, e := range s.effects {
		if e == effect {
			s.effects = append(s.effects[:i], s.effects[i+1:]...)
			break
		}
	}
}

// DecrementEffectDurations decreases the duration of all effects of a specific DurationType by 1.
func (s *Stats) DecrementEffectDurations(durationType common.DurationType) {
	for i := len(s.effects) - 1; i >= 0; i-- {
		effect := s.effects[i]
		if effect.DurationType() == durationType {
			effect.DecreaseDuration()
			if effect.Expired() {
				s.effects[i] = s.effects[len(s.effects)-1]
				s.effects = s.effects[:len(s.effects)-1]
			}
		}
	}
}

// StatEffect represents temporary or permanent modifications to stats from various sources.
type StatEffect struct {
	name        string
	description string

	agility  int
	cunning  int
	spirit   int
	strength int
	lore     int
	luck     int

	durationType common.DurationType
	duration     int
}

// Name returns the name of the StatEffect.
func (se *StatEffect) Name() string {
	return se.name
}

// Description returns the description of the StatEffect.
func (se *StatEffect) Description() string {
	return se.description
}

// DurationType returns the type of duration for the StatEffect.
func (se *StatEffect) DurationType() common.DurationType {
	return se.durationType
}

// Duration returns the remaining duration of the StatEffect.
func (se *StatEffect) Duration() int {
	return se.duration
}

// DecreaseDuration decreases the duration of the StatEffect by 1, not going below 0.
func (se *StatEffect) DecreaseDuration() {
	if se.duration > 0 {
		se.duration--
	}
}

// Expired checks if the StatEffect has expired (duration is 0).
func (se *StatEffect) Expired() bool {
	return se.duration == 0
}

// EffectedStats returns the stats modified by this effect.
func (se *StatEffect) EffectedStats() Stats {
	return Stats{
		agility:  se.agility,
		cunning:  se.cunning,
		spirit:   se.spirit,
		strength: se.strength,
		lore:     se.lore,
		luck:     se.luck,
	}
}
