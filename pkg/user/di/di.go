package di

import (
	"log"

	"github.com/Shakezidin/config"
	"github.com/Shakezidin/pkg/user/handler"
	"github.com/Shakezidin/pkg/user/repository"
	"github.com/Shakezidin/pkg/user/server"
	"github.com/Shakezidin/pkg/user/service"
	"github.com/Shakezidin/pkg/db"
)

func Init() {
	cnfg := config.LoadConfig()
	redis := config.ConnectToRedis(cnfg)
	twilio := config.SetupTwilio(cnfg)
	db := db.Database(cnfg)
	coordinatorepo := repository.NewUserRepo(db)
	coordinatorService := service.NewUserSVC(coordinatorepo, twilio, redis, cnfg)
	coordinatorHandler := handler.NewUserHandler(coordinatorService)
	err := server.NewCoordinatorGrpcServer(cnfg, coordinatorHandler)
	if err != nil {
		log.Fatalf("something went wrong", err)
	}
}
