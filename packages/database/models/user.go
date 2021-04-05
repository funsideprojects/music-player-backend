package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID          primitive.ObjectID `bson:"_id, omitempty"`
	DisplayName string             `bson:"display_name, omitempty"`
	Email       string             `bson:"email, omitempty"`
}
