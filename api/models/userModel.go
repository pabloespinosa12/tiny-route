package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID            primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Username      string             `json:"username" bson:"username"`
	Password      string             `json:"password" bson:"password"`
	Email         string             `json:"email" bson:"email"`                 // For password recovery or notifications
	CreatedAt     time.Time          `json:"createdAt" bson:"createdAt"`         // Timestamp for when the user account was created
	UpdatedAt     time.Time          `json:"updatedAt" bson:"updatedAt"`         // Timestamp for the last update to the user account
	IsVerified    bool               `json:"isVerified" bson:"isVerified"`       // To verify if the user’s email is confirmed
	ShortURLCount int                `json:"shortUrlCount" bson:"shortUrlCount"` // For tracking the number of URLs created by the user
}
