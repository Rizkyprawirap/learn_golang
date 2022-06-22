package app

import (
	"flag"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/rizkyprawirap/Toko/app/controller"
)

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func Run() {
	var server = controller.Server{}
	var appConfig = controller.AppConfig{}
	var dbConfig = controller.DBConfig{}

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	appConfig.AppName = getEnv("APP_NAME", "Toko")
	appConfig.AppEnv = getEnv("APP_ENV", "development")
	appConfig.AppPort = getEnv("APP_PORT", "8777")

	dbConfig.DBHost = getEnv("DB_HOST", "localhost")
	dbConfig.DBUser = getEnv("DB_USER", "user")
	dbConfig.DBPassword = getEnv("DB_PASSWORD", "password")
	dbConfig.DBName = getEnv("DB_NAME", "tokokoto")
	dbConfig.DBPort = getEnv("DB_PORT", "5432")

	flag.Parse()
	arg := flag.Arg(0)

	if arg != "" {
		server.InitCommands(appConfig, dbConfig)
	} else {
		server.Initialize(appConfig, dbConfig)
		server.Run(": " + appConfig.AppPort)
	}

}
