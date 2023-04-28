package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/skvenkat/golang-gin-rest-api/handlers"
)

func Routes(r *gin.Engine, g *handlers.	GoApp) {
	router := r.Use(gin.Logger(), gin.Recovery())

	router.GET("/", g.Home())

	// set up for soring details as cookies
	cookieData := cookie.NewStore([]byte("go-app"))
	router.use(sessions.Sessions("session", cookieData))
	
	router.POST("/sign-up", g.SignUp)
	router.POST("/sign-in", g.SingIn)

	authRouter := r.Group("/auth", Authorization()) 
	{
		authRouter.GET("/dashboard")
	}
}