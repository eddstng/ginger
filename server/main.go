package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"server/db"
	"server/router"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load(".env")
	fmt.Println("Starting server...")

	databaseURL := os.Getenv("DATABASE_URL")

	var err error
	for i := 0; i < 10; i++ {
		err = db.InitDBClientFromURL(databaseURL)
		if err == nil {
			break
		}
		fmt.Printf("DB connection failed, retrying in 3s... (%d/10)\n", i+1)
		time.Sleep(3 * time.Second)
	}

	if err != nil {
		log.Fatalf("Error: failed to connect to DB after retries: %v\n", err)
	}
	defer db.CloseDBClient()

	r := router.InitializeChiRouter()
	fmt.Println("Server started on http://0.0.0.0:3000")
	http.ListenAndServe("0.0.0.0:3000", r)
}
