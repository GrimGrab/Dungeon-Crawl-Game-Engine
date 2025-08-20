package engine

import (
	"SoB/internal/entity/entity"
)

type Phase string

type State struct {
	Round int            `json:"round"` // Current round number
	Turn  *entity.Entity `json:"turn"`  // Current entities turn
	Phase Phase          `json:"Phase"` // Current Phase of the game
}

func NewState() *State {
	return &State{
		Round: 0,
		Turn:  nil,
		Phase: StateGameSetup,
	}
}

func (s *State) IncRound() {
	s.Round = s.Round + 1
}

func (s *State) SetPhase(phase Phase) {
	s.Phase = phase
}

func (s *State) Reset() {
	s.Round = 0
	s.Turn = nil
	s.Phase = StateGameSetup
}

func (s *State) SetTurn(turn *entity.Entity) {
	s.Turn = turn
}

const (
	// Game Setup and Mission States
	StateGameSetup     Phase = "Game Setup"
	StateMissionSelect Phase = "Mission Select"
	StateMissionSetup  Phase = "Mission Setup"
	StateDeployHeroes  Phase = "Deploy Heroes"

	// Turn Structure
	StateStartGameTurn Phase = "Start Game Turn"
	StateEndGameTurn   Phase = "End Game Turn"
	StateStartHeroTurn Phase = "Start Hero Turn"
	StateEndHeroTurn   Phase = "End Hero Turn"

	// Hold Back the Darkness
	StateStartHoldBackDarkness Phase = "Hold Back Darkness"
	StateEndHoldBackDarkness   Phase = "End Hold Back Darkness"

	// Exploration States
	StateStartPlaceExploration  Phase = "Start Place Exploration"
	StateEndPlaceExploration    Phase = "End Place Exploration"
	StateStartRevealExploration Phase = "Start Reveal Exploration"
	StateEndRevealExploration   Phase = "End Reveal Exploration"
	StateDrawMapTile            Phase = "Draw Map Tile"
	StateConnectTile            Phase = "Connect Tile"
	StateSpawnEnemies           Phase = "Spawn Enemies"
	StateCheckSpecialRules      Phase = "Check Special Rules"
	StateEncounterCheck         Phase = "Encounter Check"
	StateDiscoverLoot           Phase = "Discover Loot"

	// Hero Action States
	StateHeroActivation    Phase = "Hero Activation"
	StateEndHeroActivation Phase = "End Hero Activation"
	StateHeroMovement      Phase = "Hero Movement"
	StateEndHeroMovement   Phase = "End Hero Movement"
	StateHeroAttack        Phase = "Hero Attack"
	StateEndHeroAttack     Phase = "End Hero Attack"
	StateHeroAbility       Phase = "Hero Ability"
	StateEndHeroAbility    Phase = "End Hero Ability"
	StateHeroScavenge      Phase = "Hero Scavenge"
	StateEndHeroScavenge   Phase = "End Hero Scavenge"
	StateHeroSearch        Phase = "Hero Search"
	StateEndHeroSearch     Phase = "End Hero Search"
	StateHeroRest          Phase = "Hero Rest"
	StateEndHeroRest       Phase = "End Hero Rest"

	// Combat States
	StateStartCombat     Phase = "Start Combat"
	StateInitiativeRoll  Phase = "Initiative Roll"
	StateCombatRound     Phase = "Combat Round"
	StateTargetSelection Phase = "Target Selection"
	StateRollToHit       Phase = "Roll To Hit"
	StateRollDamage      Phase = "Roll LoseHealth"
	StateApplyDamage     Phase = "Apply LoseHealth"
	StateEndCombat       Phase = "End Combat"

	// Injury and Death System
	StateInjuryCheck     Phase = "Injury Check"
	StateApplyInjury     Phase = "Apply Injury"
	StateDeathSave       Phase = "Death Save"
	StateCharacterDown   Phase = "Character Down"
	StateCharacterRevive Phase = "Character Revive"

	// Enemy AI States
	StateEnemyActivation Phase = "Enemy Activation"
	StateEnemyMovement   Phase = "Enemy Movement"
	StateEnemyAttack     Phase = "Enemy Attack"
	StateEnemyAbility    Phase = "Enemy Ability"
	StateEnemyEndTurn    Phase = "Enemy End Turn"

	// Growing Dread System
	StateAddGrowingDread    Phase = "Add Growing Dread"
	StateRevealGrowingDread Phase = "Reveal Growing Dread"
	StateRemoveGrowingDread Phase = "Remove Growing Dread"

	// Corruption and Mutation System
	StateCorruptionCheck  Phase = "Corruption Check"
	StateGainMutation     Phase = "Gain Mutation"
	StateCorruptionEffect Phase = "Corruption Effect"

	// Environmental Hazards
	StateHazardCheck   Phase = "Hazard Check"
	StateWeatherEffect Phase = "Weather Effect"
	StateTrapTrigger   Phase = "Trap Trigger"

	// Events and Encounters
	StateDrawEvent      Phase = "Draw Event"
	StateResolveEvent   Phase = "Resolve Event"
	StateStoryEncounter Phase = "Story Encounter"

	// Loot and Items
	StateLootRoll      Phase = "Loot Roll"
	StateDrawLoot      Phase = "Draw Loot"
	StateEquipItem     Phase = "Equip Item"
	StateUseConsumable Phase = "Use Consumable"
	StateDropItem      Phase = "Drop Item"

	// Town Phase States
	StateReturnToTown      Phase = "Return To Town"
	StateVisitSaloon       Phase = "Visit Saloon"
	StateVisitGeneralStore Phase = "Visit General Store"
	StateVisitGunsmith     Phase = "Visit Gunsmith"
	StateVisitBank         Phase = "Visit Bank"
	StateUpgradeHero       Phase = "Upgrade Hero"
	StateBuyEquipment      Phase = "Buy Equipment"
	StateSellLoot          Phase = "Sell Loot"
	StateHealInjuries      Phase = "GainHealth Injuries"

	// Mission Completion
	StateMissionComplete  Phase = "Mission Complete"
	StateMissionFailed    Phase = "Mission Failed"
	StateCalculateRewards Phase = "Calculate Rewards"
)
