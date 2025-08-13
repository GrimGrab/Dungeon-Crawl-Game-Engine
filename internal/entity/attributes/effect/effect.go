package attributes

import (
	"SoB/internal/common"
	"SoB/internal/dice"
)

type Effect struct {
	Dice
	Attribute
	name        string
	description string
	id          string

	durationType common.DurationType
	duration     int
}

func NewEffect(name, description, id string, durationType common.DurationType, duration int, dice Dice, attr Attribute) *Effect {
	return &Effect{
		name:         name,
		description:  description,
		id:           id,
		durationType: durationType,
		duration:     duration,
		Dice:         dice,
		Attribute:    attr,
	}
}

func (e *Effect) Name() string {
	return e.name
}

func (e *Effect) Description() string {
	return e.description
}

func (e *Effect) ID() string {
	return e.id
}

func (e *Effect) DurationType() common.DurationType {
	return e.durationType
}

func (e *Effect) Duration() int {
	return e.duration
}

func (e *Effect) DecrementDuration() {
	if e.duration > 0 {
		e.duration--
	}
}

func (e *Effect) IsExpired() bool {
	return e.duration == 0 && e.durationType != common.DurationTypePermanent
}

type Dice struct {
	action      common.Action // Which action it modifies
	diceAdded   []dice.Die
	diceRemoved []dice.Die
	modifier    func(results dice.RollResult) dice.RollResult
}

type Attribute struct {
	agility  int
	cunning  int
	spirit   int
	strength int
	lore     int
	luck     int

	grit   int
	sanity int
	health int

	rangedToHit int
	meleeToHit  int
	combat      int
}

func (a *Attribute) GetValue(attrType common.AttributeType) int {
	switch attrType {
	case common.AttributeAgility:
		return a.agility
	case common.AttributeCunning:
		return a.cunning
	case common.AttributeSpirit:
		return a.spirit
	case common.AttributeStrength:
		return a.strength
	case common.AttributeLore:
		return a.lore
	case common.AttributeLuck:
		return a.luck
	case common.AttributeGrit:
		return a.grit
	case common.AttributeSanity:
		return a.sanity
	case common.AttributeHealth:
		return a.health
	case common.AttributeRangedToHit:
		return a.rangedToHit
	case common.AttributeMeleeToHit:
		return a.meleeToHit
	case common.AttributeCombat:
		return a.combat
	default:
		return 0
	}
}
