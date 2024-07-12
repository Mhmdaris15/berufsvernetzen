//go:build wireinject
// +build wireinject

package main

import (
	"context"

	"github.com/google/wire"
	"gitlab.com/Mhmdaris15/berufsvernetzen/berufsvernetzen-backend/internal/configs"
	"gitlab.com/Mhmdaris15/berufsvernetzen/berufsvernetzen-backend/internal/handlers"
	"gitlab.com/Mhmdaris15/berufsvernetzen/berufsvernetzen-backend/internal/repositories"
	"gitlab.com/Mhmdaris15/berufsvernetzen/berufsvernetzen-backend/internal/services"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InitUserRepository(ctx context.Context, dbName string) (repositories.UserRepository, func(), error) {
	wire.Build(repositories.NewUserRepository, provideMongoClient, provideDatabase)
	return repositories.UserRepository{}, nil, nil
}

func InitUserService(userRepo repositories.UserRepository) services.UserService {
	wire.Build(services.NewUserService)
	return services.UserService{}
}

func InitUserHandler(userService services.UserService) handlers.UserHandler {
	wire.Build(handlers.NewUserHandler)
	return handlers.UserHandler{}
}

func provideMongoClient(ctx context.Context) (*mongo.Client, func(), error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(configs.EnvMongoURI()))
	return client, func() { client.Disconnect(ctx) }, err
}

func provideDatabase(client *mongo.Client) *mongo.Database {
	return client.Database(configs.EnvDatabaseName())
}
