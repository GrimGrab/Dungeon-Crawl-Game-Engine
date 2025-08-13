package attributes

import "SoB/internal/common"

const (
	MaxWillpower = 6
	MinWillpower = 1
)

type Willpower struct {
	willpower int

	effects []*WillpowerEffects
}

func NewWillpower(willpower int) *Willpower {
	return &Willpower{
		willpower: willpower,
	}
}

func (w *Willpower) BaseWillpower() int {
	return w.willpower
}

func (w *Willpower) TotalWillpower() int {
	total := w.willpower
	for _, effect := range w.effects {
		total += effect.WillpowerChange()
	}
	if total > MaxWillpower {
		total = MaxWillpower
	}
	if total < MinWillpower {
		total = MinWillpower
	}
	return total
}

func (w *Willpower) Effects() []*WillpowerEffects {
	return w.effects
}

func (w *Willpower) AddEffect(effect *WillpowerEffects) {
	w.effects = append(w.effects, effect)
}

func (w *Willpower) RemoveEffect(effect *WillpowerEffects) {
	for i, e := range w.effects {
		if e == effect {
			w.effects = append(w.effects[:i], w.effects[i+1:]...)
			break
		}
	}
}

func (w *Willpower) DecrementEffectDurations(durationType common.DurationType) {
	for i := len(w.effects) - 1; i >= 0; i-- {
		effect := w.effects[i]
		if effect.DurationType() == durationType {
			effect.DecreaseDuration()
			if effect.Expired() {
				w.effects[i] = w.effects[len(w.effects)-1]
				w.effects = w.effects[:len(w.effects)-1]
			}
		}
	}
}

type WillpowerEffects struct {
	name        string
	description string

	willpowerChange int

	durationType common.DurationType
	duration     int
}

func NewWillpowerEffects(name, description string, willpowerChange int, durationType common.DurationType, duration int) *WillpowerEffects {
	return &WillpowerEffects{
		name:            name,
		description:     description,
		willpowerChange: willpowerChange,
		durationType:    durationType,
		duration:        duration,
	}
}

func (we *WillpowerEffects) Name() string {
	return we.name
}

func (we *WillpowerEffects) Description() string {
	return we.description
}

func (we *WillpowerEffects) WillpowerChange() int {
	return we.willpowerChange
}

func (we *WillpowerEffects) DurationType() common.DurationType {
	return we.durationType
}

func (we *WillpowerEffects) Duration() int {
	return we.duration
}

func (we *WillpowerEffects) DecreaseDuration() {
	if we.duration > 0 {
		we.duration--
	}
}

func (we *WillpowerEffects) Expired() bool {
	return we.duration == 0
}
