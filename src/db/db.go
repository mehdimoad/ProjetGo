package db

import (
    "log"

    "github.com/mehdimoad/ProjetGo/src/models"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)
//connexion db
func Init() *gorm.DB {
    dbURL := "postgres://mehdi:pass@localhost:5432/db_go"

    db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

    if err != nil {
        log.Fatalln(err)
    }
// Migration des tables Profil et Game
    db.AutoMigrate(&models.Profil{})
    db.AutoMigrate(&models.Game{})
    return db
}