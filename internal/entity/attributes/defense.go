package attributes

import "SoB/internal/common"

const MaxDefense = 6
const MinDefense = 1

type Defense struct {
	defense int
}

func NewDefense(defense int) *Defense {
	return &Defense{
		defense: defense,
	}
}

func (d *Defense) BaseDefense() int {
	return d.defense
}

type DefenseEffects struct {
	name        string
	description string

	defense int

	durationType common.DurationType
	duration     int
}
