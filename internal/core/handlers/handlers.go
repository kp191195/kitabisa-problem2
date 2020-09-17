package handlers

import (
	"core-api/internal/core/contracts"
	clog "core-api/internal/logs"
	"encoding/json"
	"net/http"

	"github.com/spf13/viper"

	"core-api/internal/reswrapper"

	"github.com/sirupsen/logrus"
)

var svc *contracts.Services
var logger *logrus.Logger
var jwtKey []byte

func Init(s *contracts.Services) {
	logger = clog.NewLogger()
	svc = s
	jwtKey = []byte(viper.GetString("app.jwt_secret"))
}

func respondWithJson(w http.ResponseWriter, res reswrapper.ResponseWrapper) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(res.StatusCode)

	response, _ := json.Marshal(res)

	w.Write(response)
	return
}

func respondOK(w http.ResponseWriter) {
	response := reswrapper.OK()
	respondWithJson(w, response)
	return
}

func respondOKWithData(w http.ResponseWriter, payload interface{}) {
	response := reswrapper.OK()
	response.Data = payload
	respondWithJson(w, response)
	return
}

func respondInternalServerError(w http.ResponseWriter, err error) {
	response := reswrapper.InternalServerError(err)
	respondWithJson(w, response)
	return
}

func respondErrorValidation(w http.ResponseWriter, message interface{}) {
	response := reswrapper.ErrorInputValidation()
	response.Error = message
	respondWithJson(w, response)
	return
}
