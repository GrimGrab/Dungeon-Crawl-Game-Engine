package attributes

import "SoB/internal/common"

type EffectManager struct {
	effects []*Effect
}

func NewEffectManager() *EffectManager {
	return &EffectManager{
		effects: []*Effect{},
	}
}

func (em *EffectManager) AddEffect(effect *Effect) {
	em.effects = append(em.effects, effect)
}

func (em *EffectManager) RemoveEffect(effect *Effect) {
	for i, e := range em.effects {
		if e == effect {
			em.effects = append(em.effects[:i], em.effects[i+1:]...)
			return
		}
	}
}

func (em *EffectManager) Effects() []*Effect {
	return em.effects
}

func (em *EffectManager) DecrementEffects(durationType common.DurationType) {
	// Iterate backwards to avoid index shifting issues when removing elements
	for i := len(em.effects) - 1; i >= 0; i-- {
		effect := em.effects[i]
		if effect.durationType == durationType {
			effect.DecrementDuration()
			if effect.IsExpired() {
				em.RemoveEffect(effect)
			}
		}
	}
}

func (em *EffectManager) DiceModifier(action common.Action) []*Effect {
	var relevantEffects []*Effect
	for _, effect := range em.effects {
		if effect.Dice.action == action {
			relevantEffects = append(relevantEffects, effect)
		}
	}
	return relevantEffects
}

func (em *EffectManager) AttributeModifier(attribute common.AttributeType) int {
	var totalModifier int
	for _, effect := range em.effects {
		totalModifier += effect.Attribute.GetValue(attribute)
	}
	return totalModifier
}
