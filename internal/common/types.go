package common

type DurationType string
type AttributeType string
type Action string

const (
	DurationTypeCombat    DurationType = "combat"
	DurationTypeRound     DurationType = "round"
	DurationTypeAdventure DurationType = "adventure_duration"
	DurationTypePermanent DurationType = "permanent"
	DurationTypeTown      DurationType = "town"
)

const (
	AttributeAgility     AttributeType = "agility"
	AttributeCunning     AttributeType = "cunning"
	AttributeSpirit      AttributeType = "spirit"
	AttributeStrength    AttributeType = "strength"
	AttributeLore        AttributeType = "lore"
	AttributeLuck        AttributeType = "luck"
	AttributeDefense     AttributeType = "defense"
	AttributeWillpower   AttributeType = "willpower"
	AttributeHealth      AttributeType = "health"
	AttributeSanity      AttributeType = "sanity"
	AttributeGrit        AttributeType = "grit"
	AttributeRangedToHit AttributeType = "ranged_to_hit"
	AttributeMeleeToHit  AttributeType = "melee_to_hit"
	AttributeCombat      AttributeType = "combat"
)

const (
	ActionMove          Action = "move"
	ActionAttack        Action = "attack"
	ActionDefenceSave   Action = "defence_save"
	ActionWillpowerSave Action = "willpower_save"
	ActionScavenge      Action = "scavenge"
	ActionLoot          Action = "loot"
	ActionUseItem       Action = "use_item"
	ActionUseSkill      Action = "use_skill"
	ActionUseToken      Action = "use_token"
)
