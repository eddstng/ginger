package main

import (
	"fmt"
	"log"
	"server/db"
	"server/handlers"

	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Starting server...")
	err = db.InitDBClient()
	if err != nil {
		log.Fatalf("Error: %v\n", err)
	}
	defer db.CloseDBClient()

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to the server!"))
	})
	r.Get("/items", handlers.GetItemsHandler)

	fmt.Println("Server started on http://localhost:3000")
	http.ListenAndServe(":3000", r)
}
