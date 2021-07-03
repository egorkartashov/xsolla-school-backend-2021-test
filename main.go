package main

import (
	"fmt"
	"github.com/egorkartashov/xsolla-school-backend-2021-test/app"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"net/http"
	"os"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		fmt.Print(err)
	}

	dbHost := os.Getenv("POSTGRES_HOST")
	dbName := os.Getenv("POSTGRES_DB")
	dbUsername := os.Getenv("POSTGRES_USER")
	dbPassword := os.Getenv("POSTGRES_PASSWORD")
	host := fmt.Sprintf("host=%s user=%s password=%s dbname=%s", dbHost, dbUsername, dbPassword, dbName)
	db, err := gorm.Open(postgres.Open(host), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to database: %s", err.Error()))
	}

	fmt.Println("Hello, world.")

	appInstance := app.New(db)

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
