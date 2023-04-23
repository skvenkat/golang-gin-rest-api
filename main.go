package main

import (
	"context"
	"log"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/joho/godotenv"
	"github.com/skvenkat/golang-gin-rest-api/driver"
	"github.com/skvenkat/golang-gin-rest-api/modules/config"
)

var app config.GoAppTools

func main() {

	InfoLogger := log.New(os.Stdout, " ", log.LstdFlags|log.Lshortfile)
	ErrorLogger := log.New(os.Stdout, " ", log.LstdFlags|log.Lshortfile)
	app.InfoLogger = InfoLogger
	app.ErrorLogger = ErrorLogger

	err := godotenv.Load()
	if err != nil {
		app.ErrorLogger.Fatal("No .env file available")
	}

	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		app.ErrorLogger.Fatalln("mongodb uri string not found : ")
	}

	// connecting to the database
	client := driver.Connection(uri)
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			app.ErrorLogger.Fatal(err)
			return
		}
	}()

	appRouter := gin.New()
	appRouter.GET("/", func(ctx *gin.Context){
		log.Println("Creating a scalable web application with gin")
	})

	err = appRouter.Run()
	if err != nil {
		log.Fatal(err)
	}
}