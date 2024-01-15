package di

import (
	"fmt"
	"log"

	"github.com/Shakezidin/config"
	"github.com/Shakezidin/pkg/db"
	clinet "github.com/Shakezidin/pkg/user/client"
	"github.com/Shakezidin/pkg/user/handler"
	"github.com/Shakezidin/pkg/user/repository"
	"github.com/Shakezidin/pkg/user/server"
	"github.com/Shakezidin/pkg/user/service"
)

func Init() {
	cnfg := config.LoadConfig()
	redis := config.ConnectToRedis(cnfg)
	twilio := config.SetupTwilio(cnfg)
	db := db.Database(cnfg)
	clinet, err := clinet.ClientDial(*cnfg)
	if err != nil {
		fmt.Println("client dial error")
		return
	}

	coordinatorepo := repository.NewUserRepo(db)
	coordinatorService := service.NewUserSVC(coordinatorepo, twilio, redis, cnfg, clinet)
	coordinatorHandler := handler.NewUserHandler(coordinatorService)
	err = server.NewCoordinatorGrpcServer(cnfg, coordinatorHandler)
	if err != nil {
		log.Fatalf("something went wrong", err)
	}
}
