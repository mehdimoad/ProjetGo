package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mehdimoad/ProjetGo/src/handler"
	"github.com/mehdimoad/ProjetGo/src/db"
)

func main() {
	router := mux.NewRouter()
	DB:=db.Init()
	h:=handlers.New(DB)
	router.HandleFunc("/profils", h.GetAllProfils).Methods(http.MethodGet)
	router.HandleFunc("/profils/{id}", h.GetProfil).Methods(http.MethodGet)
	router.HandleFunc("/profils", h.AddProfil).Methods(http.MethodPost)
	router.HandleFunc("/profils/{id}", h.UpdateProfil).Methods(http.MethodPut)
	router.HandleFunc("/profils/{id}", h.DeleteProfil).Methods(http.MethodDelete)
	//--------------------------------------------------------------------------------
	router.HandleFunc("/games", h.GetAllGames).Methods(http.MethodGet)
	router.HandleFunc("/games/{id}", h.GetGame).Methods(http.MethodGet)
	router.HandleFunc("/games", h.AddGame).Methods(http.MethodPost)
	router.HandleFunc("/games/{id}", h.UpdateGame).Methods(http.MethodPut)
	router.HandleFunc("/games/{id}", h.DeleteGame).Methods(http.MethodDelete)

	log.Println("API is running!")
	http.ListenAndServe(":8080", router)
}