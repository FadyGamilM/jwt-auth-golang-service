package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID           primitive.ObjectID `bson:"_id"`
	FirstName    string             `json:"fname" validate:"required, min=2 max=50"`
	LastName     string             `json:"lname" validate:"required, min=2 max=50"`
	Email        string             `json:"email" validate:"email, required"`
	Password     string             `json:"password" validate:"required, min=8"`
	CreatedAt    time.Time          `json:"created_at"`
	UpdatedAt    time.Time          `json:"updated_at"`
	Token        string             `json:"token"`
	RefreshToken string             `json:"refresh_token"`
	Role         string             `json:"role" validate:"required, eq=ADMIN|eq=USER"`
}
