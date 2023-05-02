package query

import (
	"context"
	"errors"
	"fmt"
	"regexp"
	"time"

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

func (g *GoAppDB) InsertUser(user *model.User) (bool, int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 100 * time.Second)
	defer cancel()

	regMail := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	ok := regMail.MatchString(user.Email)

	if !ok {
		errMsg := fmt.Sprintf("Invalid email id provided for registration: %s", user.Email)
		g.App.ErrorLogger()
		return false, 0, errors.New(errMsg)
	}

	return false, 0, errors.New("unidentified error.")
}
