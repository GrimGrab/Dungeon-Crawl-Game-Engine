package attributes

type Sanity struct {
	maxSanity int
	sanity    int
}

func NewSanity(maxSanity int) *Sanity {
	return &Sanity{
		maxSanity: maxSanity,
		sanity:    maxSanity,
	}
}

func (s *Sanity) BaseMaxSanity() int {
	return s.maxSanity
}

func (s *Sanity) Sanity() int {
	return s.sanity
}

func (s *Sanity) LoseSanity(amount int) {
	if amount < 0 {
		return
	}
	s.sanity -= amount
	if s.sanity < 0 {
		s.sanity = 0
	}
}

func (s *Sanity) IsInsane() bool {
	return s.sanity <= 0
}
