package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Host               string
	User               string
	Password           string
	Database           string
	Port               string
	Sslmode            string
	GRPCUSERPORT       string
	SID                string
	TOKEN              string
	SERVICETOKEN       string
	SECRETKEY          string
	REDISHOST          string
	GRPCCORDINATORPORT string
	REDISPassword      string
}

func LoadConfig() *Config {
	godotenv.Load("../../.env")

	var config Config

	// Use os.Getenv to retrieve environment variables
	config.Host = os.Getenv("HOST")
	config.User = os.Getenv("USER")
	config.Password = os.Getenv("PASSWORD")
	config.Database = os.Getenv("DATABASE")
	config.Port = os.Getenv("PORT")
	config.Sslmode = os.Getenv("SSLMODE")
	config.GRPCUSERPORT = os.Getenv("GRPCUSERPORT")
	config.SID = os.Getenv("SID")
	config.TOKEN = os.Getenv("TOKEN")
	config.SERVICETOKEN = os.Getenv("SERVICETOKEN")
	config.SECRETKEY = os.Getenv("SECRETKEY")
	config.REDISHOST = os.Getenv("REDISHOST")
	config.GRPCCORDINATORPORT = os.Getenv("GRPCCORDINATORPORT")
	config.REDISPassword = os.Getenv("REDISPASSWORD")

	return &config
}
