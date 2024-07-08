package entity

const (
	AddPlayerNATSHandler = "AddPlayer"
	AddTeamNATSHandler   = "AddTeam"
	AddStatNATSHandler   = "AddStat"
	GetPlayerNATSHandler = "GetPlayer"
	GetTeamNATSHandler   = "GetTeam"
)

type Player struct {
	Name   string `json:"name"`
	TeamId int    `json:"team_id"`
}

type Team struct {
	Name string `json:"name"`
}

type Stat struct {
	PlayerID      int     `json:"player_id"`
	Points        int     `json:"points"`
	Rebounds      int     `json:"rebounds"`
	Assists       int     `json:"assists"`
	Steals        int     `json:"steals"`
	Blocks        int     `json:"blocks"`
	Fouls         int     `json:"fouls"`
	Turnovers     int     `json:"turnovers"`
	MinutesPlayed float32 `json:"minutes_played"`
}

type PlayerAverage struct {
	PlayerID      int     `json:"player_id"`
	Points        float32 `json:"points"`
	Rebounds      float32 `json:"rebounds"`
	Assists       float32 `json:"assists"`
	Steals        float32 `json:"steals"`
	Blocks        float32 `json:"blocks"`
	Fouls         float32 `json:"fouls"`
	Turnovers     float32 `json:"turnovers"`
	MinutesPlayed float32 `json:"minutes_played"`
}

type TeamAverage struct {
	TeamID        int     `json:"team_id"`
	Points        float32 `json:"points"`
	Rebounds      float32 `json:"rebounds"`
	Assists       float32 `json:"assists"`
	Steals        float32 `json:"steals"`
	Blocks        float32 `json:"blocks"`
	Fouls         float32 `json:"fouls"`
	Turnovers     float32 `json:"turnovers"`
	MinutesPlayed float32 `json:"minutes_played"`
}
