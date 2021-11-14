package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Device struct {
	ID       primitive.ObjectID `json:"id" bson:"_id"`
	Name     string             `json:"name" bson:"name"`
	Interval int                `json:"interval" bson:"interval"`
	Value    float64            `json:"value" bson:"value"`
}
