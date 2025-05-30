package router

import (
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
	r.Get("/items", handlers.GetItemsHandler())
	r.Post("/items", handlers.PostItemHandler())
	r.Put("/items", handlers.PutItemHandler())
	r.Get("/customers", handlers.GetCustomersHandler())
	r.Post("/customers", handlers.PostCustomerHandler())
	r.Put("/customers", handlers.PutCustomerHandler())
	r.Get("/orders", handlers.GetOrdersHandler())
	r.Post("/orders", handlers.PostOrderHandler())
	r.Put("/orders", handlers.PutOrderHandler())
	r.Get("/order_items", handlers.GetOrderItemsHandler())

	return r
}
