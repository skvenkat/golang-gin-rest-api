package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/skvenkat/golang-gin-rest-api/database"
	"github.com/skvenkat/golang-gin-rest-api/database/query"
	"github.com/skvenkat/golang-gin-rest-api/modules/config"
)

type GoApp struct {
	App config.GoAppTools
	DB database.DBRepo
}

func NewGoApp(app config.GoAppTools, db mongo.Client) *GoApp {
	return &GoApp{
		App: app,
		DB: query.NewGoAppDB(app, db),
	}
}

func (ga GoApp) Home() gin.HandlerFunc {
	return func(ctx gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"resp": "Welcome to gin-gonic web app"})
	}
}