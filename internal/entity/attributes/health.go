package attributes

type Health struct {
	maxHealth int
	health    int
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
