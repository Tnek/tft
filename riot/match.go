package riot

type Match struct {
	Metadata Metadata `json:"metadata"`
	Info     Info     `json:"info"`
}

type Metadata struct {
	Version string `json:"data_version"`
	ID      string `json:"match_id"`

	// Participants is a list of encrypted PUUIDs
	Participants []string `json:"participants"`
}

type Info struct {
	// Time is a unix timestamp
	Time int64 `json:"game_datetime"`

	// Length is the length of the game in seconds
	Length float64 `json:"game_length"`

	// Variation is the game variation key. As of set 3, these are Galaxy names.
	Variation string `json:"game_variation"`

	// Version is the patch that the game was played on.
	Version string `json:"game_version"`

	// Participants is game data about participants in the match. Note that this is
	// different from the Participants slice in the Metadata struct, which contains
	// PUUIDs.
	Participants []Participant `json:"participants"`

	// QueueID is the type of match it was (normal, ranked, etc).
	// See https://developer.riotgames.com/docs/lol#general_game-constants
	QueueID int `json:"queue_id"`

	Set int `json:"tft_set_number"`
}

// Participant is game data on a participant in the match.
type Participant struct {
	GoldLeft             int        `json:"gold_left"`
	LastRound            int        `json:"last_round"`
	Level                int        `json:"level"`
	Placement            int        `json:"placement"`
	PlayersEliminated    int        `json:"players_eliminated"`
	PUUID                string     `json:"puuid"`
	TimeEliminated       float64    `json:"time_eliminated"`
	TotalDamageToPlayers int        `json:"total_damage_to_players"`
	Traits               []TraitDTO `json:"traits"`
	Units                []UnitDTO  `json:"units"`
}

// TraitDTO is the object used to represent a Trait during a match.
type TraitDTO struct {
	Name        string `json:"name"`
	NumUnits    int    `json:"num_units"`
	TierCurrent int    `json:"tier_current"`
}

// UnitDTO is the object used to represent a unit during a match.
type UnitDTO struct {
	// Items is a list of item IDs. See https://developer.riotgames.com/docs/lol#data-dragon_items
	Items []int `json:"items"`

	// ID is the character ID introduced in patch 9.22 with data_version 2.
	ID string `json:"character_id"`

	Name   string `json:"name"`
	Rarity int    `json:"rarity"`
	Tier   int    `json:"tier"`
}
