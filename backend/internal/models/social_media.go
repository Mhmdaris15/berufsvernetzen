package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type SocialMedia struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Linkedin  string             `bson:"linkedin,omitempty" json:"linkedin,omitempty"`
	Facebook  string             `bson:"facebook,omitempty" json:"facebook,omitempty"`
	Instagram string             `bson:"instagram,omitempty" json:"instagram,omitempty"`
	Github    string             `bson:"github,omitempty" json:"github,omitempty"`
	Twitter   string             `bson:"twitter,omitempty" json:"twitter,omitempty"`
	Discord   string             `bson:"discord,omitempty" json:"discord,omitempty"`
}
