package models

type Profil struct {
	Id     int    `json:"id" gorm:"primaryKey"`
	Pseudo  string `json:"pseudo"`
	Rank string `json:"rank"`
	// GameId int `json:"game"`
	Games      []Game `gorm:"ForeignKey:profil_id"`
	
}