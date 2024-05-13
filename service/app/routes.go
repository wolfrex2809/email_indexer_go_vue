package app

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/wolfrex2809/email_indexer_go_vue/handler"
)

func loadRoutes() *chi.Mux {
	router := chi.NewRouter()

	router.Use(middleware.Logger)

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	router.Route("/emails", LoadEmailsRoutes)

	return router
}

func LoadEmailsRoutes(router chi.Router) {
	emailHandler := &handler.Email{}

	router.Get("/", emailHandler.List)

}
