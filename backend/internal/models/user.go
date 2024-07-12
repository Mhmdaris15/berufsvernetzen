package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Name           string             `bson:"name" json:"name,omitempty"`
	Username       string             `bson:"username" json:"username,omitempty"`
	Email          string             `bson:"email" json:"email,omitempty"`
	Gender         string             `bson:"gender" json:"gender,omitempty"`
	WhatsappNumber string             `bson:"whatsapp_number" json:"whatsapp_number,omitempty"`
	Password       string             `bson:"password" json:"password,omitempty"`
	NIK            string             `bson:"nik" json:"nik,omitempty"`
	Address        string             `bson:"address" json:"address,omitempty"`
	YearGraduation string             `bson:"year_graduation" json:"year_graduation,omitempty"`
	Birthday       string             `bson:"birthday" json:"birthday,omitempty"`
	Major          string             `bson:"major" json:"major,omitempty"`
	Languages      []string           `bson:"languages" json:"languages,omitempty"`
	Experiences    []string           `bson:"experiences" json:"experiences,omitempty"`
	SocialMedia    string             `bson:"social_media" json:"social_media,omitempty"`
	Role           string             `bson:"role" json:"role,omitempty"`
	Certifications []string           `bson:"certifications" json:"certifications,omitempty"`
	Photo          string             `bson:"photo" json:"photo,omitempty"`
	// Add other fields as per your schema.
}
