package common

type DurationType string
type AttributeType string
type CombatAttributeType string
type Action string

type Coordinates struct {
	X, Y int
}

const (
	DurationTypeCombat    DurationType = "combat"
	DurationTypeRound     DurationType = "round"
	DurationTypeAdventure DurationType = "adventure_duration"
	DurationTypePermanent DurationType = "permanent"
	DurationTypeTown      DurationType = "town"
)

const (
	AttributeAgility  AttributeType = "agility"
	AttributeCunning  AttributeType = "cunning"
	AttributeSpirit   AttributeType = "spirit"
	AttributeStrength AttributeType = "strength"
	AttributeLore     AttributeType = "lore"
	AttributeLuck     AttributeType = "luck"
	AttributeHealth   AttributeType = "health"
	AttributeSanity   AttributeType = "sanity"
	AttributeGrit     AttributeType = "grit"
)

const (
	CombatAttributeCombat      CombatAttributeType = "combat"
	CombatAttributeRangedToHit CombatAttributeType = "ranged_to_hit"
	CombatAttributeMeleeToHit  CombatAttributeType = "melee_to_hit"
	CombatAttributeDefense     CombatAttributeType = "defense"
	CombatAttributeWillpower   CombatAttributeType = "willpower"
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
