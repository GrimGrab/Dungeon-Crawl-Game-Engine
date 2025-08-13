package attributes

import "SoB/internal/common"

type Health struct {
	maxHealth int
	health    int

	effects []*HealthEffect
}

func NewHealth(maxHealth int) *Health {
	return &Health{
		maxHealth: maxHealth,
		health:    maxHealth,
	}
}

func (h *Health) BaseMaxHealth() int {
	return h.maxHealth
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

func (h *Health) RegainHealth(amount int) {
	if amount < 0 {
		return
	}
	h.health += amount
	if h.health > h.maxHealth {
		h.health = h.maxHealth
	}
}

func (h *Health) IsDead() bool {
	return h.health <= 0
}

func (h *Health) Reset() {
	h.health = h.maxHealth
}

func (h *Health) Effects() []*HealthEffect {
	return h.effects
}

func (h *Health) TotalMaxHealth() int {
	total := h.maxHealth
	for _, effect := range h.effects {
		total += effect.MaxHealthChange()
	}
	if total < 1 {
		total = 1
	}
	return total
}

func (h *Health) AddEffect(effect *HealthEffect) {
	h.effects = append(h.effects, effect)
}

func (h *Health) RemoveEffect(effect *HealthEffect) {
	for i, e := range h.effects {
		if e == effect {
			h.effects = append(h.effects[:i], h.effects[i+1:]...)
			break
		}
	}
}

func (h *Health) DecrementEffectDurations(durationType common.DurationType) {
	for i := len(h.effects) - 1; i >= 0; i-- {
		effect := h.effects[i]
		if effect.DurationType() == durationType {
			effect.DecreaseDuration()
			if effect.Expired() {
				h.effects[i] = h.effects[len(h.effects)-1]
				h.effects = h.effects[:len(h.effects)-1]
			}
		}
	}
}

type HealthEffect struct {
	name        string
	description string

	maxHealth    int
	durationType common.DurationType
	duration     int
}

func (he *HealthEffect) Name() string {
	return he.name
}

func (he *HealthEffect) Description() string {
	return he.description
}

func (he *HealthEffect) MaxHealthChange() int {
	return he.maxHealth
}

func (he *HealthEffect) DurationType() common.DurationType {
	return he.durationType
}

func (he *HealthEffect) Duration() int {
	return he.duration
}

func (he *HealthEffect) DecreaseDuration() {
	if he.duration > 0 {
		he.duration--
	}
}

func (he *HealthEffect) Expired() bool {
	return he.duration == 0
}
