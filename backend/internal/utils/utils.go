package utils

import (
	"context"
	"math/rand"

	"github.com/o1egl/paseto"
	"gitlab.com/Mhmdaris15/berufsvernetzen/berufsvernetzen-backend/internal/models"
	"gitlab.com/Mhmdaris15/berufsvernetzen/berufsvernetzen-backend/internal/pkg/database/mongodb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

const (
	SECRET_KEY = "secret"
)

var (
	maker = paseto.NewV2()
)

func GenerateRandomNumeric(length int) string {
	digits := "0123456789"
	NIK := ""
	for i := 0; i < length; i++ {
		NIK += string(digits[rand.Intn(len(digits))])
	}
	return NIK
}

func GenerateRandomString(length int) string {
	chars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	randomString := ""
	for i := 0; i < length; i++ {
		randomString += string(chars[rand.Intn(len(chars))])
	}
	return randomString
}

func GeneratePasetoToken(username string) (string, error) {
	payload := map[string]interface{}{
		"username": username,
	}

	token, err := maker.Encrypt([]byte(SECRET_KEY), payload, nil)

	if err != nil {
		return "", err
	}

	return token, nil
}

func VerifyPassword(selector, value, password string) (models.User, error) {
	var result models.User

	collection := mongodb.GetCollection(mongodb.DB, "Users")
	if collection == nil {
		return result, mongo.ErrNoDocuments
	}

	filter := bson.M{selector: value}
	err := collection.FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return result, err
		}
		return result, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(password))
	if err != nil {
		return result, err
	}

	return result, nil
}
