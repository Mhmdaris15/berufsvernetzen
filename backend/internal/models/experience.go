package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Experience struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Position    string             `bson:"position,omitempty" json:"position,omitempty"`
	Company     string             `bson:"company,omitempty" json:"company,omitempty"`
	StartDate   string             `bson:"start_date,omitempty" json:"start_date,omitempty"`
	EndDate     string             `bson:"end_date,omitempty" json:"end_date,omitempty"`
	Description string             `bson:"description,omitempty" json:"description,omitempty"`
}
