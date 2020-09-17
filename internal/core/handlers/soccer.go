package handlers

import (
	"core-api/internal/core/entities"
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

func GetSoccerTeamList(w http.ResponseWriter, r *http.Request) {
	logger.Info("Endpoint hit : GetSoccerTeamList")

	result, err := svc.Soccer.GetSoccerTeamList()
	if err != nil {
		logger.Error("Error at invoking service : Get Soccer Team List ", err)
		respondInternalServerError(w, err)
		return
	}

	respondOKWithData(w, result)
	return
}

func GetSoccerTeamDetail(w http.ResponseWriter, r *http.Request) {
	logger.Info("Endpoint hit : GetSoccerTeamDetail")

	params := mux.Vars(r)

	strTeamID := params["id"]

	teamID, err := strconv.ParseInt(strTeamID, 10, 64)
	if err != nil {
		logger.Error("Error at parsing url parameter id ", err)
		respondErrorValidation(w, "Error at parsing url parameter id")
		return
	}

	result, err := svc.Soccer.GetSoccerTeamDetail(teamID)
	if err != nil {
		if err == sql.ErrNoRows {
			respondErrorValidation(w, "Data not found!")
			return
		}
		logger.Error("Error at invoking service : GetSoccerTeamDetail ", err)
		respondInternalServerError(w, err)
		return
	}

	respondOKWithData(w, result)
	return
}

func AddSoccerTeam(w http.ResponseWriter, r *http.Request) {
	logger.Info("Endpoint hit : AddSoccerTeam")
	var input entities.FormAddSoccerTeam

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logger.Error("Error at reading request body ", err)
		respondInternalServerError(w, err)
		return
	}
	err = json.Unmarshal(body, &input)
	if err != nil {
		logger.Error("Error at parse body ", err)
		respondInternalServerError(w, err)
		return
	}

	err = svc.Soccer.AddSoccerTeam(input)
	if err != nil {
		logger.Error("Error at invoking service : Add Soccer Team", err)
		respondInternalServerError(w, err)
		return
	}
	//logger.Info(fmt.Sprintf("%v", input))

	respondOK(w)
	return
}

func AddSoccerPlayer(w http.ResponseWriter, r *http.Request) {
	logger.Info("Endpoint hit : AddSoccerPlayer")

	var input entities.FormAddSoccerPlayer

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logger.Error("Error at reading request body ", err)
		respondInternalServerError(w, err)
		return
	}
	err = json.Unmarshal(body, &input)
	if err != nil {
		logger.Error("Error at parse body ", err)
		respondInternalServerError(w, err)
		return
	}
	input.DOB, err = time.Parse("2006-01-02", input.StringDOB)
	if err != nil {
		logger.Error("Error at parse dob to date, format dob is : YYYY-MM-DD", err)
		respondErrorValidation(w, "Error at parse dob to date, format dob is : YYYY-MM-DD")
		return
	}
	err = svc.Soccer.AddSoccerPlayer(input)
	if err != nil {
		logger.Error("Error at invoking service : Add Soccer Player", err)
		respondInternalServerError(w, err)
		return
	}

	respondOK(w)
	return
}

/*func RemoveFinAccount(w http.ResponseWriter, r *http.Request) {
	logger.Info("Endpoint hit : RemoveFinAccountList")

	userSession, err := getUserSession(r)

	if err != nil {
		logger.Error("Error at gettingUserSession")
		respondErrorValidation(w, err)
		return
	}
	userID := userSession.UserSession.ID

	params := mux.Vars(r)

	accID, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		logger.Error(fmt.Sprintf("Error at parsing url parameter, id : %s", params["id"]))
		respondErrorValidation(w, fmt.Sprintf("Error at parsing to integer url parameter, id : %s", params["id"]))
		return
	}

	err = svc.FinAcc.RemoveFinAccount(accID, userID)
	if err != nil {
		logger.Error("Error at invoking service : RemoveFinAccount")
		respondInternalServerError(w, err)
		return
	}

	respondOK(w)
	return
}

func ModifyFinAccount(w http.ResponseWriter, r *http.Request) {
	logger.Info("Endpoint hit : ModifyFinAccountList")
	var input entities.FormModifyFinAccount
	userSession, err := getUserSession(r)

	if err != nil {
		logger.Error("Error at gettingUserSession")
		respondErrorValidation(w, err)
		return
	}
	userID := userSession.UserSession.ID

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logger.Error("Error at reading request body ", err)
		respondInternalServerError(w, err)
	}
	err = json.Unmarshal(body, &input)
	if err != nil {
		logger.Error("Error at parse body ", err)
		respondInternalServerError(w, err)
	}

	input.UserID = userID

	err = svc.FinAcc.ModifyFinAccount(input)
	if err != nil {
		logger.Error("Error at invoking service : ModifyFinAccount")
		respondInternalServerError(w, err)
		return
	}

	respondOK(w)
	return
}*/
