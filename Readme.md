# ProjetGo

API CRUD en GO :
Début d'un back afin d'afficher l'historique des parties d'un Profil pour un jeu.

# Bibliothèque

Gorilla/mux 
Gorm

# Base de donnée

Postgresql

## architecture
```
├── src
│    ├── db
│    │   ├── db.go
│    ├── handler
│    │   ├── handler.go
│    │
│    └── models
│        ├── game.go
│        └── profil.go
├── docker-compose.yml
├── go.mod
├── go.sum
└── main.go
```

## Initialisation

Lancer la commande pour docker
docker compose up -d

Lancer le projet
go run main.go


### Routes :

Profil :
- `GET /profils`
- `GET /profils/:id`
- `POST /profils {pseudo, rank, Games[]}`
- `PUT/profils/:id {pseudo, rank}`
- `DELETE /profils/:id`

 Game :
- `GET /games`
- `GET /games/:id`
- `POST /games {type, profil}`
- `PUT /games/:id {type, profil}`
- `DELETE /clients/:id`


## Modeles 

Profil :
- Id
- Pseudo
- Rank
- Game[]

Game :
- Id
- Type
- ProfilId