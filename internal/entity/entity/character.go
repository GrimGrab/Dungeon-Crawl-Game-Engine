package entity

import (
	"SoB/internal/entity/attributes"
	effects "SoB/internal/entity/attributes/effect"
)

type Gender string

const (
	Male   Gender = "Male"
	Female Gender = "Female"

	minimumStatValue = 1
	minimumToHitRoll = 1
	maximumToHitRoll = 6
	minimumCombat    = 1
	minimumHealth    = 1
	minimumSanity    = 1
	minimumGrit      = 0
	initialGrit      = 1
)

type Character struct {
	Entity
	sanity *attributes.Sanity
	stats  *attributes.Stats
	grit   *attributes.Grit
	gender Gender
}

type CharacterCreationOptions struct {
	Name        string
	Description string
	Gender      Gender
	Keywords    *attributes.Keywords

	MaxHealth int
	MaxSanity int
	MaxGrit   int

	RangedToHitRoll int
	MeleeToHitRoll  int
	Combat          int

	Agility  int
	Cunning  int
	Spirit   int
	Strength int
	Lore     int
	Luck     int
}

func ValidateCharacterOptions(options CharacterCreationOptions) error {
	if options.Name == "" {
		return ErrInvalidName()
	}
	if options.MaxHealth <= minimumHealth {
		return ErrInvalidHealth()
	}
	if options.MaxSanity < minimumSanity {
		return ErrInvalidSanity()
	}
	if options.MaxGrit < minimumGrit {
		return ErrInvalidGrit()
	}
	if options.RangedToHitRoll < minimumToHitRoll || options.MeleeToHitRoll < minimumToHitRoll || options.RangedToHitRoll > maximumToHitRoll || options.MeleeToHitRoll > maximumToHitRoll {
		return ErrInvalidToHitRoll()
	}
	if options.Combat < minimumCombat {
		return ErrInvalidCombat()
	}
	if options.Agility < minimumStatValue || options.Cunning < minimumStatValue || options.Spirit < minimumStatValue || options.Strength < minimumStatValue || options.Lore < minimumStatValue || options.Luck < minimumStatValue {
		return ErrInvalidStats()
	}
	return nil
}

func NewCharacter(id string, options CharacterCreationOptions) (*Character, error) {
	effectManager := effects.NewEffectManager()
	return &Character{
		Entity: Entity{
			id:             id,
			name:           options.Name,
			description:    options.Description,
			kind:           KindCharacter,
			health:         attributes.NewHealth(options.MaxHealth, effectManager),
			offensiveStats: attributes.NewOffensiveStats(options.RangedToHitRoll, options.MeleeToHitRoll, options.Combat, effectManager),
			keywords:       options.Keywords,
			effectManager:  effectManager,
		},
		sanity: attributes.NewSanity(options.MaxSanity, effectManager),
		stats:  attributes.NewStats(options.Agility, options.Cunning, options.Spirit, options.Strength, options.Lore, options.Luck, effectManager),
		grit:   attributes.NewGrit(initialGrit, options.MaxGrit, effectManager),
		gender: options.Gender,
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

func (c *Character) Gender() Gender {
	return c.gender
}
