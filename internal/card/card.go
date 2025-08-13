package card

type Card interface {
	Name() string
	Description() string
	Type() string
	ID() string
}
