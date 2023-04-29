package main

import (
	"encoding/gob"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/skvenkat/golang-gin-rest-api/driver"
	"github.com/skvenkat/golang-gin-rest-api/handlers"
	"github.com/skvenkat/golang-gin-rest-api/modules/config"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var app config.GoAppTools
var validate *validator.Validate

func main() {
	gob.Register(map[string]interface{}{})
	gob.Register(primitive.NewObjectID())

		InfoLogger := log.New(os.Stdout, " ", log.LstdFlags|log.Lshortfile)
		ErrorLogger := log.New(os.Stdout, " ", log.LstdFlags|log.Lshortfile)

		validate = validator.New()
		app.InfoLogger = InfoLogger
		app.ErrorLogger = ErrorLogger
		app.Validate = validate

		err := godotenv.Load()
		if err != nil {
			app.ErrorLogger.Fatal("No .env file available to load the config")
		}

		uri := os.Getenv("MONGODB_URI")
		if uri == "" {
			app.ErrorLogger.Fatalln("mongodb uri env var is not set :")	
		}

		client := driver.Connection(uri)
		defer func() {
			if err = client.Disconnect(); err != nil {
				app.ErrorLogger.Fatal("DB connection couldn't be closed: \n", err)
			}
		}()

		appRouter := gin.New()

		goApp := handlers.NewGoApp(&app, client)
		Routes(appRouter, goApp)

		err = appRouter.Run()
		if err != nil {
			app.ErrorLogger.Fatal("Failed to start the app server:\n", err)
		}

}