package common

type DurationType string

const (
	DurationTypeCombat    DurationType = "combat"
	DurationTypeRound     DurationType = "round"
	DurationTypeAdventure DurationType = "adventure_duration"
	DurationTypePermanent DurationType = "permanent"
	DurationTypeTown      DurationType = "town"
)
