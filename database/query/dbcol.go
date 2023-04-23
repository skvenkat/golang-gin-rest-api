package query

import "go.mongodb.org/mongo-driver/mongo"

func User(db mongo.Client, collection string) mongo.Collection {
	var user = db.Database("gin_rest_api").Collection(collection)
	return user
}
