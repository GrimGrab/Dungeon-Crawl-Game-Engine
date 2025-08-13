package attributes

import (
	"SoB/internal/common"
	effects "SoB/internal/entity/attributes/effect"
)

type Sanity struct {
	maxSanity int
	sanity    int

	effectManager *effects.EffectManager
}

func NewSanity(maxSanity int, effectManager *effects.EffectManager) *Sanity {
	return &Sanity{
		maxSanity: maxSanity,
		sanity:    maxSanity,

		effectManager: effectManager,
	}
}

func (s *Sanity) BaseMaxSanity() int {
	return s.maxSanity
}

func (s *Sanity) MaxSanity() int {
	return s.maxSanity + s.effectManager.AttributeModifier(common.AttributeSanity)
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

func (s *Sanity) GainSanity(amount int) {
	if amount < 0 {
		return
	}
	s.sanity += amount
	if s.sanity > s.maxSanity+s.effectManager.AttributeModifier(common.AttributeSanity) {
		s.sanity = s.maxSanity + s.effectManager.AttributeModifier(common.AttributeSanity)
	}
}

func (s *Sanity) IsInsane() bool {
	return s.sanity <= 0
}
