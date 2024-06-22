package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"server/db"
	"server/router"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Starting server...")
	databaseURL := os.Getenv("DATABASE_URL")
	err = db.InitDBClient(databaseURL)
	if err != nil {
		log.Fatalf("Error: %v\n", err)
	}
	defer db.CloseDBClient()
	r := router.InitializeChiRouter()
	fmt.Println("Server started on http://localhost:3000")
	http.ListenAndServe(":3000", r)
}
