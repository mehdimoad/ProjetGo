package db

import (
    "log"

    "github.com/mehdimoad/ProjetGo/pkg/models"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

func Init() *gorm.DB {
    dbURL := "postgres://mehdi:pass@localhost:5432/db_go"

    db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

    if err != nil {
        log.Fatalln(err)
    }

    db.AutoMigrate(&models.Profil{})
    db.AutoMigrate(&models.Game{})
    return db
}