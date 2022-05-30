package models

type Game struct {
	Id     int    `json:"id" gorm:primaryKey`
	Type  string `json:"type"`
	ProfilId int `json:"profil"`
	// Profils []Profil `gorm:"ForeignKey:game_id"`
	// Winner []int `json:"winner"`
	// Looser []int `json:"looser"`
}
