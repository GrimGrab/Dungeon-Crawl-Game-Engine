package attributes

import "SoB/internal/common"

const MaxDefense = 6
const MinDefense = 1

type Defense struct {
	defense int

	effects []*DefenseEffects
}

func NewDefense(defense int) *Defense {
	return &Defense{
		defense: defense,
	}
}

func (d *Defense) BaseDefense() int {
	return d.defense
}

func (d *Defense) TotalDefense() int {
	total := d.defense
	for _, effect := range d.effects {
		total += effect.DefenseChange()
	}
	if total > MaxDefense {
		total = MaxDefense
	}
	if total < MinDefense {
		total = MinDefense
	}
	return total
}

func (d *Defense) Effects() []*DefenseEffects {
	return d.effects
}

func (d *Defense) AddEffect(effect *DefenseEffects) {
	d.effects = append(d.effects, effect)
}

func (d *Defense) RemoveEffect(effect *DefenseEffects) {
	for i, e := range d.effects {
		if e == effect {
			d.effects = append(d.effects[:i], d.effects[i+1:]...)
			break
		}
	}
}

func (d *Defense) DecrementEffectDurations(durationType common.DurationType) {
	for i := len(d.effects) - 1; i >= 0; i-- {
		effect := d.effects[i]
		if effect.DurationType() == durationType {
			effect.DecreaseDuration()
			if effect.Expired() {
				d.effects[i] = d.effects[len(d.effects)-1]
				d.effects = d.effects[:len(d.effects)-1]
			}
		}
	}
}

type DefenseEffects struct {
	name        string
	description string

	defense int

	durationType common.DurationType
	duration     int
}

func NewDefenseEffects(name, description string, defense int, durationType common.DurationType, duration int) *DefenseEffects {
	return &DefenseEffects{
		name:         name,
		description:  description,
		defense:      defense,
		durationType: durationType,
		duration:     duration,
	}
}

func (de *DefenseEffects) Name() string {
	return de.name
}

func (de *DefenseEffects) Description() string {
	return de.description
}

func (de *DefenseEffects) DefenseChange() int {
	return de.defense
}

func (de *DefenseEffects) DurationType() common.DurationType {
	return de.durationType
}

func (de *DefenseEffects) Duration() int {
	return de.duration
}

func (de *DefenseEffects) DecreaseDuration() {
	if de.duration > 0 {
		de.duration--
	}
}

func (de *DefenseEffects) Expired() bool {
	return de.duration == 0
}
