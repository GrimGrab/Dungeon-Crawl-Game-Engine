package entity

import (
	"SoB/internal/entity/attributes"
	effects "SoB/internal/entity/attributes/effect"
)

type Kind string

const (
	KindCharacter Kind = "Character"
	KindMonster   Kind = "Monster"
)

type Entity struct {
	id          string
	name        string
	description string
	kind        Kind

	health         *attributes.Health
	offensiveStats *attributes.OffensiveStats
	keywords       *attributes.Keywords
	effectManager  *effects.EffectManager
}

func (e *Entity) ID() string {
	return e.id
}

func (e *Entity) Name() string {
	return e.name
}

func (e *Entity) Description() string {
	return e.description
}

func (e *Entity) Type() Kind {
	return e.kind
}

func (e *Entity) Health() *attributes.Health {
	return e.health
}

func (e *Entity) Keywords() *attributes.Keywords {
	return e.keywords
}

func (e *Entity) OffensiveStats() *attributes.OffensiveStats {
	return e.offensiveStats
}

func (e *Entity) EffectManager() *effects.EffectManager {
	return e.effectManager
}
