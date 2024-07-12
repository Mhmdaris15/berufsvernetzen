package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Certification struct {
	ID           primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name         string             `bson:"name,omitempty" json:"name,omitempty"`
	Organization string             `bson:"organization,omitempty" json:"organization,omitempty"`
	StartDate    string             `bson:"start_date,omitempty" json:"start_date,omitempty"`
	EndDate      string             `bson:"end_date,omitempty" json:"end_date,omitempty"`
	Description  string             `bson:"description,omitempty" json:"description,omitempty"`
}
