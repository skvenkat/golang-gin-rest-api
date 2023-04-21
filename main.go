package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	appRouter := gin.New()
	appRouter.GET("/", func(ctx *gin.Context){
		log.Println("Creating a scalable web application with gin")
	})

	err := appRouter.Run()
	if err != nil {
		log.Fatal(err)
	}
}