package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RSVPData struct {
	DisplayName string                       `json:"displayName" binding:"required"`
	RSVP        string                       `json:"rsvpOptions" binding:"required"`
	People      map[string]map[string]string `bson:"people" json:"people" binding:"required"`
}

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
