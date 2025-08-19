package card

type Deck interface {
	Draw() (Card, error)
	Shuffle()
	Insert(card Card)
	DrawX(x int) ([]Card, error)
	LookAtTop(x int) []Card
	Count() int
}
