package attributes

import (
	"SoB/internal/common"
	effects "SoB/internal/entity/attributes/effect"
)

type Health struct {
	maxHealth int
	health    int

	effectManager *effects.EffectManager
}

func NewHealth(maxHealth int, effectsManager *effects.EffectManager) *Health {
	return &Health{
		maxHealth:     maxHealth,
		health:        maxHealth,
		effectManager: effectsManager,
	}
}

func (h *Health) BaseMaxHealth() int {
	return h.maxHealth
}

func (h *Health) MaxHealth() int {
	return h.maxHealth + h.effectManager.AttributeModifier(common.AttributeHealth)
}

func (h *Health) Health() int {
	return h.health
}

func (h *Health) LoseHealth(amount int) {
	if amount < 0 {
		return
	}
	h.health -= amount
	if h.health < 0 {
		h.health = 0
	}
}

func (h *Health) GainHealth(amount int) {
	if amount < 0 {
		return
	}
	h.health += amount
	if h.health > h.maxHealth+h.effectManager.AttributeModifier(common.AttributeHealth) {
		h.health = h.maxHealth + h.effectManager.AttributeModifier(common.AttributeHealth)
	}
}

func (h *Health) IsDead() bool {
	return h.health <= 0
}

func (h *Health) Reset() {
	h.health = h.maxHealth + h.effectManager.AttributeModifier(common.AttributeHealth)
}
