package repositories

import (
	"context"
	"errors"
	"log"

	"gitlab.com/Mhmdaris15/berufsvernetzen/berufsvernetzen-backend/internal/configs"
	"gitlab.com/Mhmdaris15/berufsvernetzen/berufsvernetzen-backend/internal/models"
	"gitlab.com/Mhmdaris15/berufsvernetzen/berufsvernetzen-backend/internal/pkg/database/mongodb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository interface {
	GetUsers() ([]models.User, error)
	GetUser(id string) (models.User, error)
	CreateUser(user models.User) (models.User, error)
	UpdateUser(user models.User) (models.User, error)
	DeleteUser(id string) (models.User, error)
}

type UserRepositoryImpl struct {
	db *mongo.Client
	UserRepository
}

func NewUserRepository(db *mongo.Client) *UserRepositoryImpl {
	return &UserRepositoryImpl{db: db}
}

func (r *UserRepositoryImpl) GetUsers() ([]models.User, error) {
	var users []models.User

	cursor, err := r.db.Database(configs.EnvDatabaseName()).Collection("Users").Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}

	if err = cursor.All(context.Background(), &users); err != nil {
		return nil, err
	}

	return users, nil
}

func (r *UserRepositoryImpl) GetUser(id string) (models.User, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return models.User{}, err
	}

	filter := bson.M{"_id": objectID}

	var user models.User

	userDoc := r.db.Database(configs.EnvDatabaseName()).Collection("Users").FindOne(context.Background(), filter)

	err = userDoc.Decode(&user)
	if err != nil {
		return models.User{}, err
	}

	return user, err
}

func (r *UserRepositoryImpl) CreateUser(user models.User) (models.User, error) {
	log.Printf("user: %+v", user)
	// Name, Username, Email, and Password are required
	if user.Name == "" || user.Username == "" || user.Email == "" || user.Password == "" {
		return models.User{}, errors.New("name, username, email, and password are required")
	}

	insertedUser, err := r.db.Database(configs.EnvDatabaseName()).Collection("Users").InsertOne(context.Background(), user)
	if err != nil {
		return models.User{}, err
	}

	user.ID = insertedUser.InsertedID.(primitive.ObjectID)

	return user, err
}

func (r *UserRepositoryImpl) UpdateUser(user models.User) (models.User, error) {
	// Patch the user
	filter := bson.M{"_id": user.ID}

	update := bson.M{
		"$set": bson.M{},
	}

	if user.Name != "" {
		update["$set"].(bson.M)["name"] = user.Name
	}
	if user.Username != "" {
		update["$set"].(bson.M)["username"] = user.Username
	}
	if user.Email != "" {
		update["$set"].(bson.M)["email"] = user.Email
	}
	if user.Gender != "" {
		update["$set"].(bson.M)["gender"] = user.Gender
	}
	if user.WhatsappNumber != "" {
		update["$set"].(bson.M)["whatsapp_number"] = user.WhatsappNumber
	}
	if user.Password != "" {
		update["$set"].(bson.M)["password"] = user.Password
	}
	if user.NIK != "" {
		update["$set"].(bson.M)["nik"] = user.NIK
	}
	if user.Address != "" {
		update["$set"].(bson.M)["address"] = user.Address
	}
	if user.YearGraduation != "" {
		update["$set"].(bson.M)["year_graduation"] = user.YearGraduation
	}
	if user.Birthday != "" {
		update["$set"].(bson.M)["birthday"] = user.Birthday
	}
	if user.Major != "" {
		update["$set"].(bson.M)["major"] = user.Major
	}
	if user.Languages != nil {
		update["$set"].(bson.M)["languages"] = user.Languages
	}
	if user.Experiences != nil {
		update["$set"].(bson.M)["experiences"] = user.Experiences
	}
	if user.SocialMedia != "" {
		update["$set"].(bson.M)["social_media"] = user.SocialMedia
	}
	if user.Role != "" {
		update["$set"].(bson.M)["role"] = user.Role
	}
	if user.Certifications != nil {
		update["$set"].(bson.M)["certifications"] = user.Certifications
	}
	if user.Photo != "" {
		update["$set"].(bson.M)["photo"] = user.Photo
	}

	result, err := mongodb.UserCollection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return models.User{}, err
	}

	if result.MatchedCount == 0 {
		return models.User{}, errors.New("user not found")
	}

	return user, nil
}

func (r *UserRepositoryImpl) DeleteUser(id string) (models.User, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return models.User{}, err
	}

	filter := bson.M{"_id": objectID}

	var user models.User

	userDoc := r.db.Database(configs.EnvDatabaseName()).Collection("Users").FindOneAndDelete(context.Background(), filter)

	err = userDoc.Decode(&user)
	if err != nil {
		return models.User{}, err
	}

	return user, err
}
