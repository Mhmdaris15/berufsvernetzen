package configs

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

func EnvPort() string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return os.Getenv("PORT")
}

func EnvMongoURI() string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file " + err.Error())
	}
	return os.Getenv("MONGO_URI")
}

func EnvDatabaseName() string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return os.Getenv("DATABASE_NAME")
}

func EnvGRPCPort() string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return os.Getenv("GRPC_PORT")
}

func EnvPasetoSymmetricKey() string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return os.Getenv("PASETO_SYMMETRIC_KEY")
}

func EnvAccessTokenDuration() time.Duration {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	timeDuration := os.Getenv("ACCESS_TOKEN_DURATION")
	duration, err := time.ParseDuration(timeDuration)
	if err != nil {
		log.Fatal("Error parsing duration")
	}

	return duration
}

func EnvMeiliMasterKey() string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return os.Getenv("MEILISEARCH_MASTER_KEY")
}

func EnvGCloudProjectID() string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return os.Getenv("GCLOUD_PROJECT_ID")
}
