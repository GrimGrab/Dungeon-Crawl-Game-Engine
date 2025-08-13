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
		if effect.DiceEffect.action == action {
			relevantEffects = append(relevantEffects, effect)
		}
	}
	return relevantEffects
}

func (em *EffectManager) AttributeModifier(attribute common.AttributeType) int {
	var totalModifier int
	for _, effect := range em.effects {
		totalModifier += effect.Attribute.GetAttributeValue(attribute)
	}
	return totalModifier
}

func (em *EffectManager) CombatAttributeModifier(attribute common.CombatAttributeType) int {
	if attribute == common.CombatAttributeCombat {
		total := 0
		for _, effect := range em.effects {
			total += effect.combat
		}
		return total
	}

	var minModifier *int // Use pointer to distinguish between 0 and no value found

	for _, effect := range em.effects {
		// Skip effects marked as penalty
		if effect.CombatAttribute.IsForced() {
			return effect.GetCombatValue(attribute)
		}

		value := effect.CombatAttribute.GetCombatValue(attribute)

		// If this is the first non-penalty value we've found, or if it's lower than our current minimum
		if minModifier == nil || value < *minModifier {
			minModifier = &value
		}
	}

	// Return the minimum value found, or 0 if no non-penalty effects exist
	if minModifier == nil {
		return 0
	}
	return *minModifier
}
