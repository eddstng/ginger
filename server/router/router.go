package router

import (
	"server/db"
	"server/handlers"

	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func InitializeChiRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Ginger API"))
	})
	dbGetter := db.NewDBGetter()
	r.Get("/items", handlers.GetItemsHandler(dbGetter))
	return r
}
