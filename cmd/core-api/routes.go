package main

import (
	"core-api/internal/core/handlers"

	"github.com/gorilla/mux"
)

func initRoutes(router *mux.Router) {
	// Soccer API
	router.HandleFunc("/team", handlers.GetSoccerTeamList).Methods("GET")
	router.HandleFunc("/team/{id}", handlers.GetSoccerTeamDetail).Methods("GET")
	router.HandleFunc("/team", handlers.AddSoccerTeam).Methods("POST")
	router.HandleFunc("/player", handlers.AddSoccerPlayer).Methods("POST")

}
