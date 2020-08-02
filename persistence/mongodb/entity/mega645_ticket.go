package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type Mega645Ticket struct {
	Id     *primitive.ObjectID `bson:"_id,omitempty"`
	Ticket []int               `bson:"ticket"`
	Status int                 `bson:"status"`
}
