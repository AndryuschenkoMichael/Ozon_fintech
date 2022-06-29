package handler

import (
	"Ozon_fintech/pkg/service"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
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

	return router
}
