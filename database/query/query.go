package query

import (
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/skvenkat/golang-gin-rest-api/modules/config"
	"github.com/skvenkat/golang-gin-rest-api/modules/database"
)

type GoAppDB struct {
	App config.GoAppTools
	DB mongo.Client
}

func NewGoAppDB(app config.GoAppTools, db mongo.Client) database.DBRepo {
	return &GoAppDB{
		App: app,
		DB: db,
	}
}

func (g *GoAppDB) InsertUser() {
	return
}
