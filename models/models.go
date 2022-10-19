package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Guest struct {
	ID          primitive.ObjectID  `bson:"_id,omitempty"`
	DisplayName string              `bson:"displayName,omitempty"`
	Guests      []map[string]string `bson:"meals,omitempty"`
}

type Settings struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Password string             `bson:"pw,omitempty"`
}
