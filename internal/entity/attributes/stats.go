package attributes

const MinimumStatValue = 1

// Stats represents the core attributes of a character, including base values and any active effect.
type Stats struct {
	agility  int
	cunning  int
	spirit   int
	strength int
	lore     int
	luck     int
}

// NewStats creates a new Stats instance with the given base attribute values.
func NewStats(agility, cunning, spirit, strength, lore, luck int) *Stats {
	return &Stats{
		agility:  agility,
		cunning:  cunning,
		spirit:   spirit,
		strength: strength,
		lore:     lore,
		luck:     luck,
	}
}

// BaseAgility returns the base agility value without any effect.
func (s *Stats) BaseAgility() int {
	return s.agility
}

// BaseCunning returns the base cunning value without any effect.
func (s *Stats) BaseCunning() int {
	return s.cunning
}

// BaseSpirit returns the base spirit value without any effect.
func (s *Stats) BaseSpirit() int {
	return s.spirit
}

// BaseStrength returns the base strength value without any effect.
func (s *Stats) BaseStrength() int {
	return s.strength
}

// BaseLore returns the base lore value without any effect.
func (s *Stats) BaseLore() int {
	return s.lore
}

// BaseLuck returns the base luck value without any effect.
func (s *Stats) BaseLuck() int {
	return s.luck
}
