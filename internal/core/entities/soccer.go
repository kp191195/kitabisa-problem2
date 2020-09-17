package entities

import "time"

type SoccerTeam struct {
	ID               int64  `json:"id" db:"id"`
	Name             string `json:"name" db:"name"`
	CoachName        string `json:"coach_name" db:"coach_name"`
	TotalPlayer      int64  `json:"total_player" db:"total_player"`
	PlayerAgeAverage int64  `json:"player_age_average" db:"player_age_average"`
}

type SoccerPlayer struct {
	ID          int64  `json:"id" db:"id"`
	TeamID      int64  `json:"team_id" db:"team_id"`
	Fullname    string `json:"fullname" db:"fullname"`
	Age         int8   `json:"age" db:"age"`
	DOB         string `json:"dob" db:"dob"`
	Position    string `json:"position" db:"position"`
	Nationality string `json:"nationality" db:"nationality"`
	ShirtNumber string `json:"shirt_number" db:"shirt_number"`
}

type SoccerTeamDetail struct {
	SoccerTeam
	Player []SoccerPlayer `json:"player_list"`
}

type FormAddSoccerTeam struct {
	Name      string `json:"name" db:"name"`
	CoachName string `json:"coach_name" db:"coach_name"`
}

type FormAddSoccerPlayer struct {
	TeamID      int64     `json:"team_id" db:"team_id"`
	Fullname    string    `json:"fullname" db:"fullname"`
	Age         int8      `json:"age" db:"age"`
	StringDOB   string    `json:"dob" db:"-"`
	DOB         time.Time `json:"-" db:"dob"`
	Position    string    `json:"position" db:"position"`
	Nationality string    `json:"nationality" db:"nationality"`
	ShirtNumber string    `json:"shirt_number" db:"shirt_number"`
}
