package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID              primitive.ObjectID `bson:"_id"`
	Email           string             `bson:"email"`
	Username        string             `bson:"username"`
	Password        string             `bson:"password"`
	Fullname        string             `bson:"display_name"`
	ProfileImage    string             `bson:"profile_image, omitempty"`
	ProfileImageURL string             `bson:"profile_image_url, omitempty"`
	CreatedAt       time.Time          `bson:"created_at"`
	UpdatedAt       time.Time          `bson:"updated_at"`
}
