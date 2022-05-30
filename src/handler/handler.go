package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"fmt"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
	"github.com/mehdimoad/ProjetGo/src/models"
	"gorm.io/gorm"
)
type handler struct {
	DB *gorm.DB
}
func New(db *gorm.DB) handler{
	return handler{db}
}
//-----------------CRUD Profil---------------------------------------------------//
func (h handler) AddProfil(w http.ResponseWriter, r *http.Request) {
	// Lire le corp de la requete
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}
    
	var profil models.Profil
	json.Unmarshal(body, &profil)

	if result := h.DB.Create(&profil); result.Error != nil{
		fmt.Println(result.Error)
	}

	// Retourne code 201 quand le profil a ete cree
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("Created")
}
func (h handler) DeleteProfil(w http.ResponseWriter, r *http.Request) {
    // lire dynamiquement l'id en tant que parametre
    vars := mux.Vars(r)
    id, _ := strconv.Atoi(vars["id"])

    // trouver le profil par lid

    var profil models.Profil

    if result := h.DB.First(&profil, id); result.Error != nil {
        fmt.Println(result.Error)
    }

    // supprime le profil
    h.DB.Delete(&profil)

    w.Header().Add("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode("Deleted")
}

func (h handler) GetAllProfils(w http.ResponseWriter, r *http.Request) {
    var profils []models.Profil

    if result := h.DB.Preload("Games").Find(&profils); result.Error != nil {
        fmt.Println(result.Error)
    }

    w.Header().Add("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(profils)
}

func (h handler) GetProfil(w http.ResponseWriter, r *http.Request) {
    
    vars := mux.Vars(r)
    id, _ := strconv.Atoi(vars["id"])

    
    var profil models.Profil

    if result := h.DB.First(&profil, id); result.Error != nil {
        fmt.Println(result.Error)
    }

    w.Header().Add("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(profil)
}
func (h handler) UpdateProfil(w http.ResponseWriter, r *http.Request) {
    
    vars := mux.Vars(r)
    id, _ := strconv.Atoi(vars["id"])

    
    defer r.Body.Close()
    body, err := ioutil.ReadAll(r.Body)

    if err != nil {
        log.Fatalln(err)
    }

    var updatedProfil models.Profil
    json.Unmarshal(body, &updatedProfil)

    var profil models.Profil

    if result := h.DB.First(&profil, id); result.Error != nil {
        fmt.Println(result.Error)
    }

    profil.Pseudo = updatedProfil.Pseudo
    profil.Rank = updatedProfil.Rank
    

    h.DB.Save(&profil)

    w.Header().Add("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode("Updated")
}
//-----------------CRUD GAME---------------------------------------------------//
func (h handler) AddGame(w http.ResponseWriter, r *http.Request) {
	
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var game models.Game
	json.Unmarshal(body, &game)

	if result := h.DB.Create(&game); result.Error != nil{
		fmt.Println(result.Error)
	}

	
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("Created")
}
func (h handler) DeleteGame(w http.ResponseWriter, r *http.Request) {
    
    vars := mux.Vars(r)
    id, _ := strconv.Atoi(vars["id"])

   

    var game models.Game

    if result := h.DB.First(&game, id); result.Error != nil {
        fmt.Println(result.Error)
    }

   
    h.DB.Delete(&game)

    w.Header().Add("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode("Deleted")
}

func (h handler) GetAllGames(w http.ResponseWriter, r *http.Request) {
    var games []models.Game

    if result := h.DB.Find(&games); result.Error != nil {
        fmt.Println(result.Error)
    }

    w.Header().Add("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(games)
}

func (h handler) GetGame(w http.ResponseWriter, r *http.Request) {
    
    vars := mux.Vars(r)
    id, _ := strconv.Atoi(vars["id"])

    
    var game models.Game

    if result := h.DB.First(&game, id); result.Error != nil {
        fmt.Println(result.Error)
    }

    w.Header().Add("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(game)
}
func (h handler) UpdateGame(w http.ResponseWriter, r *http.Request) {
    
    vars := mux.Vars(r)
    id, _ := strconv.Atoi(vars["id"])

    
    defer r.Body.Close()
    body, err := ioutil.ReadAll(r.Body)

    if err != nil {
        log.Fatalln(err)
    }

    var updatedGame models.Game
    json.Unmarshal(body, &updatedGame)

    var game models.Game

    if result := h.DB.First(&game, id); result.Error != nil {
        fmt.Println(result.Error)
    }

    game.Type = updatedGame.Type
    

    h.DB.Save(&game)

    w.Header().Add("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode("Updated")
}