package contracts

import "core-api/internal/core/entities"

type SoccerService interface {
	GetSoccerTeamList() (result []entities.SoccerTeam, err error)
	GetSoccerTeamDetail(id int64) (result entities.SoccerTeamDetail, err error)
	AddSoccerTeam(input entities.FormAddSoccerTeam) (err error)
	AddSoccerPlayer(input entities.FormAddSoccerPlayer) (err error)
}
