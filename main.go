package main

import (
	"fmt"
	"github.com/egorkartashov/xsolla-school-backend-2021-test/app"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"net/http"
	"os"
)

const (
	postgresHostEnvKey     = "POSTGRES_HOST"
	postgresDbEnvKey       = "POSTGRES_DB"
	postgresUsernameEnvKey = "POSTGRES_USERNAME"
	postgresPasswordEnvKey = "POSTGRES_PASSWORD"
)

func main() {

	dbHost := os.Getenv(postgresHostEnvKey)
	dbName := os.Getenv(postgresDbEnvKey)
	dbUsername := os.Getenv(postgresUsernameEnvKey)
	dbPassword := os.Getenv(postgresPasswordEnvKey)
	host := fmt.Sprintf("host=%s user=%s password=%s dbname=%s", dbHost, dbUsername, dbPassword, dbName)

	db, err := gorm.Open(postgres.Open(host), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to database: %s\n", err.Error()))
	}

	logger := configureLogger()

	appInstance, err := app.New(db, logger)
	if err != nil {
		panic(fmt.Sprintf("Failed to start application: %s\n", err.Error()))
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	fmt.Println("Using port: " + port)

	err = http.ListenAndServe(":"+port, appInstance.Router)

	if err != nil {
		fmt.Print(err)
	}
}

func configureLogger() *log.Entry {
	log.SetFormatter(&log.JSONFormatter{})

	standardFields := log.Fields{
		"appname":  "products-api",
		"hostname": "localhost",
	}

	defaultLogger := log.WithFields(standardFields)

	f, err := os.OpenFile("applogs/errors", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		defaultLogger.Fatalf("Error opening file for logs: %s", err)
	}
	defer func(f *os.File) {
		_ = f.Close()
	}(f)

	log.SetOutput(f)

	return defaultLogger
}
