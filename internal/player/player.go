package player

import "SoB/internal/character"

type Player struct {
	id        string
	alias     string
	character *character.Character
}

func (p *Player) ID() string                      { return p.id }
func (p *Player) Alias() string                   { return p.alias }
func (p *Player) Character() *character.Character { return p.character }
