package app

import (
	"context"
	"fmt"
	"net/http"

	"github.com/wolfrex2809/email_indexer_go_vue/config"
)

type App struct {
	router http.Handler
}

func New() *App {
	app := &App{
		router: loadRoutes(),
	}

	return app
}

func (a *App) Start(ctx context.Context) error {

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", config.ConfigEnv.Port),
		Handler: a.router,
	}
	fmt.Println("Application is running...")

	err := server.ListenAndServe()
	if err != nil {
		return fmt.Errorf("there was an error starting the server: %w", err)
	}

	return nil
}
