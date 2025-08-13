package attributes

import "SoB/internal/common"

type Sanity struct {
	maxSanity int
	sanity    int

	effects []*SanityEffect
}

func NewSanity(maxSanity int) *Sanity {
	return &Sanity{
		maxSanity: maxSanity,
		sanity:    maxSanity,
	}
}

func (s *Sanity) BaseMaxSanity() int {
	return s.maxSanity
}

func (s *Sanity) Sanity() int {
	return s.sanity
}

func (s *Sanity) LoseSanity(amount int) {
	if amount < 0 {
		return
	}
	s.sanity -= amount
	if s.sanity < 0 {
		s.sanity = 0
	}
}

func (s *Sanity) RegainSanity(amount int) {
	if amount < 0 {
		return
	}
	s.sanity += amount
	totalMax := s.TotalMaxSanity()
	if s.sanity > totalMax {
		s.sanity = totalMax
	}
}

func (s *Sanity) IsInsane() bool {
	return s.sanity <= 0
}

func (s *Sanity) Reset() {
	s.sanity = s.TotalMaxSanity()
}

func (s *Sanity) Effects() []*SanityEffect {
	return s.effects
}

func (s *Sanity) TotalMaxSanity() int {
	total := s.maxSanity
	for _, effect := range s.effects {
		total += effect.MaxSanityChange()
	}
	if total < 1 {
		total = 1
	}
	return total
}

func (s *Sanity) AddEffect(effect *SanityEffect) {
	s.effects = append(s.effects, effect)
}

func (s *Sanity) RemoveEffect(effect *SanityEffect) {
	for i, e := range s.effects {
		if e == effect {
			s.effects = append(s.effects[:i], s.effects[i+1:]...)
			break
		}
	}
}

func (s *Sanity) DecrementEffectDurations(durationType common.DurationType) {
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

type SanityEffect struct {
	name        string
	description string

	maxSanity    int
	durationType common.DurationType
	duration     int
}

func (se *SanityEffect) Name() string {
	return se.name
}

func (se *SanityEffect) Description() string {
	return se.description
}

func (se *SanityEffect) MaxSanityChange() int {
	return se.maxSanity
}

func (se *SanityEffect) DurationType() common.DurationType {
	return se.durationType
}

func (se *SanityEffect) Duration() int {
	return se.duration
}

func (se *SanityEffect) DecreaseDuration() {
	if se.duration > 0 {
		se.duration--
	}
}

func (se *SanityEffect) Expired() bool {
	return se.duration == 0
}
