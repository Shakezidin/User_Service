package di

import (
	"log"

	"github.com/Shakezidin/config"
	"github.com/Shakezidin/pkg/db"
	client "github.com/Shakezidin/pkg/user/client"
	"github.com/Shakezidin/pkg/user/handler"
	"github.com/Shakezidin/pkg/user/repository"
	"github.com/Shakezidin/pkg/user/server"
	"github.com/Shakezidin/pkg/user/service"
)

// Init initializes the dependencies and starts the server.
func Init() {
	// Load configuration
	cfg := config.LoadConfig()

	// Connect to Redis
	redis, err := config.ConnectToRedis(cfg)
	if err != nil {
		log.Fatalf("failed to connect to Redis: %v", err)
	}

	// Setup Twilio
	twilio := config.SetupTwilio(cfg)

	// Connect to database
	dbConn := db.Database(cfg)

	// Initialize gRPC client
	grpcClient, err := client.ClientDial(*cfg)
	if err != nil {
		log.Fatalf("failed to dial gRPC client: %v", err)
	}

	// Create repository
	userRepo := repository.NewUserRepo(dbConn)

	// Create user service
	userService := service.NewUserSVC(userRepo, twilio, redis, cfg, grpcClient)

	// Create user handler
	userHandler := handler.NewUserHandler(userService)

	// Start gRPC server
	err = server.NewCoordinatorGrpcServer(cfg, userHandler)
	if err != nil {
		log.Fatalf("failed to start gRPC server: %v", err)
	}
}
