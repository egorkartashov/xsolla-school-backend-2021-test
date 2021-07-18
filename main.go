package main

import (
	"fmt"
	"github.com/egorkartashov/xsolla-school-backend-2021-test/app"
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

	appInstance, err := app.New(db)
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
