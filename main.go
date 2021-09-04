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
	postgresHostEnvKey     =  "POSTGRES_HOST"
	postgresDbEnvKey       = "POSTGRES_DB"
	postgresUsernameEnvKey = "POSTGRES_USERNAME"
	postgresPasswordEnvKey = "POSTGRES_PASSWORD"
)

func main() {

	dbHost := os.Getenv(postgresHostEnvKey)
	dbName := os.Getenv(postgresDbEnvKey)
	dbUsername := os.Getenv(postgresUsernameEnvKey)
	dbPassword := os.Getenv(postgresPasswordEnvKey)
	dbPort := os.Getenv("POSTGRES_PORT")
	if dbPort == "" {
		dbPort = "5432"
	}

	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v",
		dbHost, dbUsername, dbPassword, dbName, dbPort)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to database: %v", err))
	}

	appInstance, err := app.New(db)
	if err != nil {
		panic(fmt.Sprintf("Failed to start application: %v", err))
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
