package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Url struct {
	ID  primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Url string             `json: "url" bson:"url"`
}
