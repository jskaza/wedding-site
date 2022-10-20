package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// type Meals struct {
// 	ID          primitive.ObjectID  `bson:"_id,omitempty"`
// 	DisplayName string              `bson:"displayName,omitempty"`
// 	NGuests     int32               `bson:"nGuests,omitempty"`
// 	Meals       []map[string]string `bson:"meals,omitempty"`
// }

type Guest struct {
	ID          primitive.ObjectID  `bson:"_id,omitempty"`
	DisplayName string              `bson:"displayName,omitempty"`
	NGuests     int32               `bson:"nGuests,omitempty"`
	Meals       []map[string]string `bson:"meals,omitempty"`
}

type Settings struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Password string             `bson:"pw,omitempty"`
}
