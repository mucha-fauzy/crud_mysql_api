package routes

import (
	"crud_mysql_api/internal/services"
	"net/http"

	"crud_mysql_api/internal/handlers"

	"github.com/go-chi/chi"
)

type Router struct {
	Handler *handlers.Handler
}

func NewRouter(service services.Service) *Router {
	handler := handlers.NewHandler(service)
	return &Router{
		Handler: handler,
	}
}

func (r *Router) SetupRoutes() http.Handler {
	mux := chi.NewRouter()

	mux.Route("/api/v1", func(rc chi.Router) {
		rc.Post("/product", r.Handler.CreateProduct)
		rc.Get("/product", r.Handler.ListProducts)
		rc.Put("/variant/{variantID}", r.Handler.UpdateVariant)
		rc.Delete("/product/{productID}", r.Handler.SoftDeleteProduct)
		rc.Delete("/product/hard/{productID}", r.Handler.HardDeleteProduct)
	})
	return mux
}
