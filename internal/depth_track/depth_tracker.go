package depth_track

type DepthManager interface {
	PosseLevel() int
	DarknessLevel() int
	IncreasePosseLevel()
	DecreasePosseLevel()
	IncreaseDarknessLevel()
	DecreaseDarknessLevel()
	MinimumDarknessRoll() int
}

// DepthTracker manages the levels of Posse and Darkness in the game.
type DepthTracker struct {
	maxLevel      int
	posseLevel    int
	darknessLevel int
	minLevel      int
}

// New creates a new DepthTracker with specified min and max levels, and starting levels for Posse and Darkness.
func New(minLevel, maxLevel, posseStart, darknessStart int) (*DepthTracker, error) {
	if minLevel > maxLevel {
		return nil, ErrMinGreaterThanMax(minLevel, maxLevel)
	}
	if posseStart < minLevel || posseStart > maxLevel {
		return nil, ErrLevelOutOfBounds(posseStart, minLevel, maxLevel)
	}
	if darknessStart < minLevel || darknessStart > maxLevel {
		return nil, ErrLevelOutOfBounds(darknessStart, minLevel, maxLevel)
	}
	return &DepthTracker{
		minLevel:      minLevel,
		maxLevel:      maxLevel,
		posseLevel:    posseStart,
		darknessLevel: darknessStart,
	}, nil
}

// PosseLevel returns the current Posse level.
func (dt *DepthTracker) PosseLevel() int {
	return dt.posseLevel
}

// DarknessLevel returns the current Darkness level.
func (dt *DepthTracker) DarknessLevel() int {
	return dt.darknessLevel
}

// IncreasePosseLevel increases the Posse level by 1, up to the maximum level.
func (dt *DepthTracker) IncreasePosseLevel() {
	if dt.posseLevel < dt.maxLevel {
		dt.posseLevel++
	}
}

// DecreasePosseLevel decreases the Posse level by 1, down to the minimum level.
func (dt *DepthTracker) DecreasePosseLevel() {
	if dt.posseLevel > dt.minLevel {
		dt.posseLevel--
	}
}

// IncreaseDarknessLevel increases the Darkness level by 1, up to the maximum level.
func (dt *DepthTracker) IncreaseDarknessLevel() {
	if dt.darknessLevel < dt.maxLevel {
		dt.darknessLevel++
	}
}

// DecreaseDarknessLevel decreases the Darkness level by 1, down to the minimum level.
func (dt *DepthTracker) DecreaseDarknessLevel() {
	if dt.darknessLevel > dt.minLevel {
		dt.darknessLevel--
	}
}

// MinimumDarknessRoll returns the minimum roll required to hold back the Darkness based on the current Posse level.
func (dt *DepthTracker) MinimumDarknessRoll() int {
	if dt.PosseLevel() > 10 {
		return 7
	} else if dt.PosseLevel() > 5 {
		return 8
	} else {
		return 9
	}
}
