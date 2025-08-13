package player

import "SoB/internal/entity/entity"

type Player struct {
	id        string
	alias     string
	character *entity.Character
}

func (p *Player) ID() string                   { return p.id }
func (p *Player) Alias() string                { return p.alias }
func (p *Player) Character() *entity.Character { return p.character }
