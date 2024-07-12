package main

import (
	"log"
	"sync"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gitlab.com/Mhmdaris15/berufsvernetzen/berufsvernetzen-backend/internal/configs"
	"gitlab.com/Mhmdaris15/berufsvernetzen/berufsvernetzen-backend/internal/pkg/database/meilisearch"
	"gitlab.com/Mhmdaris15/berufsvernetzen/berufsvernetzen-backend/internal/pkg/database/mongodb"
	"gitlab.com/Mhmdaris15/berufsvernetzen/berufsvernetzen-backend/internal/routes"
	"gitlab.com/Mhmdaris15/berufsvernetzen/berufsvernetzen-backend/internal/token"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(2) // Number of Goroutines to wait for

	go func() {
		defer wg.Done()
		mongodb.ConnectDB()

		meilisearch.ConnectToMeiliSearch()
		jobIndex := meilisearch.InitializeIndex("Jobs")

		log.Printf("Job Index: %v", jobIndex)

		// deletedJobs, err := meilisearch.DeleteDocumentsFromIndex(jobIndex)
		// if err != nil {
		// 	log.Printf("Error when deleting documents from index: %s", err.Error())
		// }
		// log.Printf("Deleted Jobs from MeiliSearch: %v", deletedJobs)

		// insertedJobs, err := meilisearch.AddDocumentsToIndex(jobIndex)
		// if err != nil {
		// 	log.Printf("Error when adding documents to index: %s", err.Error())
		// }
		// log.Printf("Inserted Jobs to MeiliSearch: %v", insertedJobs)
	}()

	go func() {
		defer wg.Done()
		RunConnectGinGonic()
	}()

	wg.Wait()
}

func RunConnectGinGonic() {
	log.Print(configs.EnvPasetoSymmetricKey())
	tokenMaker, err := token.NewPasetoMaker(configs.EnvPasetoSymmetricKey())
	if err != nil {
		log.Fatalf("cannot create token maker: %s", err.Error())
	}

	token.TokenMaker = tokenMaker

	r := gin.Default()

	// Setup CORS middleware
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowHeaders = []string{"Origin", "Content-Type", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Access-Control-Allow-Methods", "Authorization", "X-Requested-With", "Accept", "Accept-Encoding", "Accept-Language", "Connection", "Host", "Origin", "Referer", "User-Agent", "Username"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"}
	config.AllowCredentials = true

	r.Use(cors.New(config))

	// Setup Router
	routes.SetupRoutes(r)

	// Run the server
	err = r.Run(":" + configs.EnvPort()) // listen and serve on 0.0.0.0:3000
	if err != nil {
		log.Printf("Error when running server: %s", err.Error())
	}
}

// func RunGRPCServer() {
// 	lis, err := net.Listen("tcp", ":"+configs.EnvGRPCPort())
// 	if err != nil {
// 		log.Fatalf("failed to listen: %v", err)
// 	}

// 	itemService := &handlers.ItemServiceServerImpl{}
// 	sayHelloService := &handlers.HelloServer{}
// 	userService := &handlers.UserServiceServerImpl{}
// 	seedService := &handlers.SeedServiceServerImpl{}

// 	s := grpc.NewServer()
// 	berufsvernetzen.RegisterItemServiceServer(s, itemService)
// 	berufsvernetzen.RegisterSayHelloServer(s, sayHelloService)
// 	berufsvernetzen.RegisterUserServiceServer(s, userService)
// 	berufsvernetzen.RegisterSeedServiceServer(s, seedService)
// 	log.Println("Starting gRPC server on port " + configs.EnvGRPCPort())
// 	if err := s.Serve(lis); err != nil {
// 		log.Fatalf("failed to serve: %v", err)
// 	}
// }
