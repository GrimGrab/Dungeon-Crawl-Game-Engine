package tokens

import "SoB/internal/engine"

type Token struct {
	name   string
	action func(p *engine.Entity) error
	price  int
}
