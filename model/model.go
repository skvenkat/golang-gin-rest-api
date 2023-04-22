package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID			primitive.ObjectID	`json:"_id" bson:"_id"`
	Name		string
	Email		string
	Age			int
	Password	string
	Occupation	string
	Token		string
	NewToken	string
	CreatedAt	time.Time
	UpdatedAt	time.Time

}