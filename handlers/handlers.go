package handlers

import (
	"errors"
	"net/http"
	"time"

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

func (ga *GoApp) SignUp() gin.HandlerFunc {
    return func (ctx *gin.Context) {
        var user *model.User
        if err := ctx.ShouldBindJSON(&user); err != nil {
            _ = ctx.AbortWithError(http.StatusBadRequest, gin.Error{Err: err})
        }
        user.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
        user.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
        user.Password, _ = encrypt.Hash(user.Password)
        if err := ga.App.Validate.Struct(&user); err != nil {
            if _, ok := err.(*validator.InvalidValidationError); !ok {
                _ = ctx.AbortWithError(http.StatusBadRequest, gin.Error{Err: err})
                ga.App.InfoLogger.Println(err)
                return
            }
        }

        ok, status, err := ga.DB.InsertUser(user)
        if err != nil {
            _ = ctx.AbortWithError(http.StatusInternalServerError, errors.New("error while adding new user"))
            ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
            return
        }
        if !ok {
            _ = ctx.AbortWithError(http.StatusInternalServerError, err)
            return
        }

        switch status {
        case 1:
            ctx.JSON(http.StatusOK, gin.H{"message": "Registered Successfully"})
            return
        case 2:
            ctx.JSON(http.StatusFound, gin.H{"message": "Existing Account, Go to the Login page"})
            return
        }
    }
}
