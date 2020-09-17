package soccer

import (
	"core-api/internal/core/contracts"
	"core-api/internal/core/entities"

	"github.com/jmoiron/sqlx"
)

type Service struct {
	db *sqlx.DB
}

func Init(app *contracts.App) contracts.SoccerService {
	svc := Service{db: app.DB}
	return &svc
}

func (svc *Service) GetSoccerTeamList() (result []entities.SoccerTeam, err error) {
	var emptyResult []entities.SoccerTeam
	err = svc.db.Select(&result, getSoccerTeamList)
	if err != nil {
		return
	}

	if len(result) == 0 {
		result = emptyResult
		return
	}
	return
}

func (svc *Service) GetSoccerTeamDetail(id int64) (result entities.SoccerTeamDetail, err error) {
	var soccerTeam entities.SoccerTeam
	var soccerPlayer []entities.SoccerPlayer
	err = svc.db.Get(&soccerTeam, getSoccerTeamById, id)
	if err != nil {
		return
	}

	err = svc.db.Select(&soccerPlayer, getSoccerPlayerByTeamId, id)
	if err != nil {
		return
	}

	if len(soccerPlayer) == 0 {
		soccerPlayer = []entities.SoccerPlayer{}
	}

	result.SoccerTeam = soccerTeam
	result.Player = soccerPlayer

	return
}

func (svc *Service) AddSoccerTeam(input entities.FormAddSoccerTeam) (err error) {
	_, err = svc.db.NamedExec(addSoccerTeam, input)
	if err != nil {
		return
	}
	return
}

func (svc *Service) AddSoccerPlayer(input entities.FormAddSoccerPlayer) (err error) {
	_, err = svc.db.NamedExec(addSoccerPlayer, input)
	if err != nil {
		return
	}
	return
}
