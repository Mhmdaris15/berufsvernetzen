package seeder

import (
	"context"
	"errors"
	"log"
	"math/rand"
	"strconv"
	"sync"

	"github.com/brianvoe/gofakeit/v6"
	"gitlab.com/Mhmdaris15/berufsvernetzen/berufsvernetzen-backend/internal/models"
	"gitlab.com/Mhmdaris15/berufsvernetzen/berufsvernetzen-backend/internal/pkg/database/mongodb"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

func SeedUser() ([]models.User, error) {
	var users []models.User
	var roles = []string{"admin", "user"}
	var genders = []string{"male", "female"}

	userCollection := mongodb.GetCollection(mongodb.DB, "Users")
	experienceCollection := mongodb.GetCollection(mongodb.DB, "Experiences")
	certificateCollection := mongodb.GetCollection(mongodb.DB, "Certificates")
	socialMediaCollection := mongodb.GetCollection(mongodb.DB, "SocialMedias")

	for i := 0; i < 10; i++ {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte("2wsx1qaz"), 16)
		if err != nil {
			log.Fatalf("Error when generating password: %s", err.Error())
		}
		user := models.User{
			ID:             primitive.NewObjectID(),
			Name:           gofakeit.Name(),
			Username:       gofakeit.Username(),
			Email:          gofakeit.Email(),
			Password:       string(hashedPassword),
			Gender:         genders[rand.Intn(len(genders))],
			WhatsappNumber: gofakeit.Phone(),
			// NIK:            utils.GenerateRandomString(16),
			NIK:            gofakeit.CreditCard().Number,
			Address:        gofakeit.Address().Address,
			Role:           roles[rand.Intn(len(roles))],
			YearGraduation: strconv.Itoa(gofakeit.Date().Year()),
			Birthday:       gofakeit.Date().String(),
			Major:          gofakeit.JobTitle(),
			Languages:      []string{gofakeit.Language()},
			Photo:          gofakeit.ImageURL(100, 100),
		}

		userExperiences := []string{}
		userCertificates := []string{}

		// Create 2 Experiences
		for j := 0; j < 2; j++ {
			newExperience := models.Experience{
				// ID:          primitive.NewObjectID(),
				Position:    gofakeit.JobTitle(),
				Company:     gofakeit.Company(),
				StartDate:   gofakeit.Date().String(),
				EndDate:     gofakeit.Date().String(),
				Description: gofakeit.Sentence(10),
			}

			newCertificate := models.Certification{
				// ID:           primitive.NewObjectID(),
				Name:         gofakeit.MovieName(),
				Organization: gofakeit.Company(),
				StartDate:    gofakeit.Date().String(),
				EndDate:      gofakeit.Date().String(),
				Description:  gofakeit.Sentence(10),
			}

			// Insert newExperience to experienceCollection
			insertedExperience, err := experienceCollection.InsertOne(context.Background(), newExperience)
			if err != nil {
				log.Fatalf("Error when inserting experience: %s", err.Error())
			}
			userExperiences = append(userExperiences, insertedExperience.InsertedID.(primitive.ObjectID).Hex())

			// Insert newCertificate to experienceCollection
			insertedCertificate, err := certificateCollection.InsertOne(context.Background(), newCertificate)
			if err != nil {
				log.Fatalf("Error when inserting certificate: %s", err.Error())
			}
			userCertificates = append(userCertificates, insertedCertificate.InsertedID.(primitive.ObjectID).Hex())
		}

		// Insert userExperiences and userCertificates to user
		user.Experiences = userExperiences
		user.Certifications = userCertificates

		// Insert SocialMedia
		newSocialMedia := models.SocialMedia{
			// ID:       primitive.NewObjectID(),
			Linkedin:  gofakeit.URL(),
			Facebook:  gofakeit.URL(),
			Twitter:   gofakeit.URL(),
			Instagram: gofakeit.URL(),
			Discord:   gofakeit.URL(),
			Github:    gofakeit.URL(),
		}
		insertedSocialMedia, err := socialMediaCollection.InsertOne(context.Background(), newSocialMedia)
		if err != nil {
			log.Fatalf("Error when inserting social media: %s", err.Error())
		}

		user.SocialMedia = insertedSocialMedia.InsertedID.(primitive.ObjectID).Hex()

		users = append(users, user)
	}

	var userInterfaces = make([]interface{}, len(users))
	for i, user := range users {
		userInterfaces[i] = user
	}

	var wg sync.WaitGroup
	insertedUsersChan := make(chan []models.User, 1) // Channel to collect results from goroutines

	// Number of goroutines to use
	numGoroutines := 4

	// Divide users into chunks for each goroutine
	chunkSize := (len(users) + numGoroutines - 1) / numGoroutines

	// Spawn multiple goroutines for concurrent insertion
	for i := 0; i < numGoroutines; i++ {
		start := i * chunkSize
		end := start + chunkSize
		if end > len(users) {
			end = len(users)
		}

		wg.Add(1)
		go func(usersSlice []interface{}) {
			defer wg.Done()

			// Check if the `id` field is set for each user before inserting it into the database.
			var usersToUpdate []interface{}
			for _, user := range usersSlice {
				u := user.(models.User)
				if u.ID == primitive.NilObjectID {
					// Generate a random UUID or use a sequence generator to set the `id` field.
					u.ID = primitive.NewObjectID()
				}
				usersToUpdate = append(usersToUpdate, u)
			}

			results, err := userCollection.InsertMany(context.Background(), usersToUpdate)
			if err != nil {
				log.Fatalf("Error when inserting users: %s", err.Error())

				// Send an empty result to the channel to indicate error
				insertedUsersChan <- []models.User{}
			}

			var insertedUsers []models.User
			for _, result := range results.InsertedIDs {
				var user models.User
				user.ID = result.(primitive.ObjectID)
				insertedUsers = append(insertedUsers, user)
			}

			// Send the result to the channel
			insertedUsersChan <- insertedUsers
		}(userInterfaces[start:end])
	}

	// Close the channel once all goroutines are done
	go func() {
		wg.Wait()
		close(insertedUsersChan)
	}()

	// Collect results from the channel
	var insertedUsers []models.User
	for result := range insertedUsersChan {
		insertedUsers = append(insertedUsers, result...)
	}

	// If the number of inserted users is null, return an error
	if len(insertedUsers) == 0 {
		// Create an error
		err := errors.New("error when seeding users")

		return nil, err
	}

	// Return users where exist in insertedUsers, otherwise return delete from users
	// for i, user := range users {
	// 	found := false
	// 	for _, insertedUser := range insertedUsers {
	// 		if user.ID == insertedUser.ID {
	// 			found = true
	// 			break
	// 		}
	// 	}
	// 	if !found {
	// 		users = append(users[:i], users[i+1:]...)
	// 	}
	// }

	// log.Print(users)

	// If the number of inserted users is not different from the number of users, log the success message
	if len(insertedUsers) == len(users) {
		log.Print("Successfully seeded users")
	}

	return users, nil
}
