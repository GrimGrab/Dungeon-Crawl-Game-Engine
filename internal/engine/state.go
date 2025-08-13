package engine

type State string

const (
	// Game Setup and Mission States
	StateGameSetup     State = "Game Setup"
	StateMissionSelect State = "Mission Select"
	StateMissionSetup  State = "Mission Setup"
	StateDeployHeroes  State = "Deploy Heroes"

	// Turn Structure
	StateStartGameTurn State = "Start Game Turn"
	StateEndGameTurn   State = "End Game Turn"
	StateStartHeroTurn State = "Start Hero Turn"
	StateEndHeroTurn   State = "End Hero Turn"

	// Hold Back the Darkness
	StateStartHoldBackDarkness State = "Hold Back Darkness"
	StateEndHoldBackDarkness   State = "End Hold Back Darkness"

	// Exploration States
	StateStartPlaceExploration  State = "Start Place Exploration"
	StateEndPlaceExploration    State = "End Place Exploration"
	StateStartRevealExploration State = "Start Reveal Exploration"
	StateEndRevealExploration   State = "End Reveal Exploration"
	StateDrawMapTile            State = "Draw Map Tile"
	StateConnectTile            State = "Connect Tile"
	StateSpawnEnemies           State = "Spawn Enemies"
	StateCheckSpecialRules      State = "Check Special Rules"
	StateEncounterCheck         State = "Encounter Check"
	StateDiscoverLoot           State = "Discover Loot"

	// Hero Action States
	StateHeroActivation    State = "Hero Activation"
	StateEndHeroActivation State = "End Hero Activation"
	StateHeroMovement      State = "Hero Movement"
	StateEndHeroMovement   State = "End Hero Movement"
	StateHeroAttack        State = "Hero Attack"
	StateEndHeroAttack     State = "End Hero Attack"
	StateHeroAbility       State = "Hero Ability"
	StateEndHeroAbility    State = "End Hero Ability"
	StateHeroScavenge      State = "Hero Scavenge"
	StateEndHeroScavenge   State = "End Hero Scavenge"
	StateHeroSearch        State = "Hero Search"
	StateEndHeroSearch     State = "End Hero Search"
	StateHeroRest          State = "Hero Rest"
	StateEndHeroRest       State = "End Hero Rest"

	// Combat States
	StateStartCombat     State = "Start Combat"
	StateInitiativeRoll  State = "Initiative Roll"
	StateCombatRound     State = "Combat Round"
	StateTargetSelection State = "Target Selection"
	StateRollToHit       State = "Roll To Hit"
	StateRollDamage      State = "Roll LoseHealth"
	StateApplyDamage     State = "Apply LoseHealth"
	StateEndCombat       State = "End Combat"

	// Injury and Death System
	StateInjuryCheck     State = "Injury Check"
	StateApplyInjury     State = "Apply Injury"
	StateDeathSave       State = "Death Save"
	StateCharacterDown   State = "Character Down"
	StateCharacterRevive State = "Character Revive"

	// Enemy AI States
	StateEnemyActivation State = "Enemy Activation"
	StateEnemyMovement   State = "Enemy Movement"
	StateEnemyAttack     State = "Enemy Attack"
	StateEnemyAbility    State = "Enemy Ability"
	StateEnemyEndTurn    State = "Enemy End Turn"

	// Growing Dread System
	StateAddGrowingDread    State = "Add Growing Dread"
	StateRevealGrowingDread State = "Reveal Growing Dread"
	StateRemoveGrowingDread State = "Remove Growing Dread"

	// Corruption and Mutation System
	StateCorruptionCheck  State = "Corruption Check"
	StateGainMutation     State = "Gain Mutation"
	StateCorruptionEffect State = "Corruption Effect"

	// Environmental Hazards
	StateHazardCheck   State = "Hazard Check"
	StateWeatherEffect State = "Weather Effect"
	StateTrapTrigger   State = "Trap Trigger"

	// Events and Encounters
	StateDrawEvent      State = "Draw Event"
	StateResolveEvent   State = "Resolve Event"
	StateStoryEncounter State = "Story Encounter"

	// Loot and Items
	StateLootRoll      State = "Loot Roll"
	StateDrawLoot      State = "Draw Loot"
	StateEquipItem     State = "Equip Item"
	StateUseConsumable State = "Use Consumable"
	StateDropItem      State = "Drop Item"

	// Town Phase States
	StateReturnToTown      State = "Return To Town"
	StateVisitSaloon       State = "Visit Saloon"
	StateVisitGeneralStore State = "Visit General Store"
	StateVisitGunsmith     State = "Visit Gunsmith"
	StateVisitBank         State = "Visit Bank"
	StateUpgradeHero       State = "Upgrade Hero"
	StateBuyEquipment      State = "Buy Equipment"
	StateSellLoot          State = "Sell Loot"
	StateHealInjuries      State = "GainHealth Injuries"

	// Mission Completion
	StateMissionComplete  State = "Mission Complete"
	StateMissionFailed    State = "Mission Failed"
	StateCalculateRewards State = "Calculate Rewards"
)
