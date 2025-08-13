package entity

import (
	"SoB/internal/entity/attributes"
)

type Gender string

const (
	Male   Gender = "Male"
	Female Gender = "Female"
)

type Character struct {
	Entity
	sanity *attributes.Sanity
	stats  *attributes.Stats
	grit   *attributes.Grit
	gender Gender
}

func NewCharacter(id, name, description string, gender Gender, health *attributes.Health, sanity *attributes.Sanity, grit *attributes.Grit, stats *attributes.Stats, offensiveStats *attributes.OffensiveStats, keywords *attributes.Keywords) (*Character, error) {
	if health == nil || sanity == nil || grit == nil || stats == nil {
		return nil, ErrInvalidEntityConstruction("health, sanity, grit, and stats must be provided when creating a character")
	}
	return &Character{
		Entity: Entity{
			id:             id,
			name:           name,
			description:    description,
			kind:           KindCharacter,
			health:         health,
			offensiveStats: offensiveStats,
			keywords:       keywords,
		},
		sanity: sanity,
		stats:  stats,
		grit:   grit,
		gender: gender,
	}, nil
}

func (c *Character) Sanity() *attributes.Sanity {
	return c.sanity
}

func (c *Character) Stats() *attributes.Stats {
	return c.stats
}

func (c *Character) Grit() *attributes.Grit {
	return c.grit
}
