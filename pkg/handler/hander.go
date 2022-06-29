package handler

import (
	_ "Ozon_fintech/docs"
	"Ozon_fintech/pkg/service"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	httpSwagger "github.com/swaggo/http-swagger"
)

type Handler struct {
	service *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{
		service: services,
	}
}

func (h *Handler) InitRoutes() chi.Router {
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	router.Route("/api", func(r chi.Router) {
		r.Get("/get-full-link", h.getFullLink)
		r.Post("/post-link", h.postLink)
	})

	router.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8000/swagger/doc.json")))

	return router
}
