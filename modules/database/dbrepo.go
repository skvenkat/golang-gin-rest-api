package database

import (
	"github.com/skvenkat/golang-gin-rest-api/modules/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DBRepo interface {
	InsertUser(user *model.User) (bool, int, error)
	VerifyUser(email string) (primitive.M, error)
	UpdateInfo(userID primitive.ObjectID, tk map[string]string) (bool, error)
}
