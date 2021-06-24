package main

import (
	"fmt"
	"github.com/egorkartashov/xsolla-school-backend-2021-test/controllers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"net/http"
	"os"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		fmt.Print(err)
	}

	fmt.Println("Hello, world.")

	router := mux.NewRouter()
	router.HandleFunc("/api/ping", controllers.GetPing)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	fmt.Println("Using port: " + port)

	err = http.ListenAndServe(":"+port, router)

	if err != nil {
		fmt.Print(err)
	}
}
