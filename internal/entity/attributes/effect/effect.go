package attributes

import (
	"SoB/internal/common"
	"SoB/internal/dice"
)

type Effect struct {
	DiceEffect
	Attribute
	CombatAttribute
	name        string
	description string
	id          string

	durationType common.DurationType
	duration     int
}

func NewEffect(name, description, id string, durationType common.DurationType, duration int, dice DiceEffect, attr Attribute) *Effect {
	return &Effect{
		name:         name,
		description:  description,
		id:           id,
		durationType: durationType,
		duration:     duration,
		DiceEffect:   dice,
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

type DiceEffect struct {
	action      common.Action // Which action it modifies
	diceAdded   []dice.Die
	diceRemoved []dice.Die
	modifier    func(results dice.RollResult) dice.RollResult
	reroll      int
}

func (de *DiceEffect) Action() common.Action {
	return de.action
}

func (de *DiceEffect) DiceAdded() []dice.Die {
	return de.diceAdded
}

func (de *DiceEffect) DiceRemoved() []dice.Die {
	return de.diceRemoved
}

func (de *DiceEffect) Modifier() func(results dice.RollResult) dice.RollResult {
	return de.modifier
}

func (de *DiceEffect) Reroll() int {
	return de.reroll
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
}

func (a *Attribute) GetAttributeValue(attrType common.AttributeType) int {
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
	default:
		return 0
	}
}

type CombatAttribute struct {
	rangedToHit int
	meleeToHit  int
	combat      int
	defense     int
	willpower   int

	force bool // There are instances where an effect forces a value, this indicates that and overrides other effects
}

func (c *CombatAttribute) IsForced() bool {
	return c.force
}

func (c *CombatAttribute) GetCombatValue(attrType common.CombatAttributeType) int {
	switch attrType {
	case common.CombatAttributeRangedToHit:
		return c.rangedToHit
	case common.CombatAttributeMeleeToHit:
		return c.meleeToHit
	case common.CombatAttributeCombat:
		return c.combat
	case common.CombatAttributeDefense:
		return c.defense
	case common.CombatAttributeWillpower:
		return c.willpower
	default:
		return 0
	}
}
