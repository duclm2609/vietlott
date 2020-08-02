package mongodb

import "go.mongodb.org/mongo-driver/mongo"

type handler struct {
	db *mongo.Database
}

func NewHandler(db *mongo.Database) handler {
	return handler{db: db}
}
