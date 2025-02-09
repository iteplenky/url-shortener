package app

import (
	"github.com/iteplenky/url-shortener/internal/app/endpoint"
	"github.com/iteplenky/url-shortener/internal/app/generator"
	"github.com/iteplenky/url-shortener/internal/app/middleware"
	"github.com/iteplenky/url-shortener/internal/app/store"
	"log"
	"net/http"
)

type App struct {
	e   *endpoint.Endpoint
	mux *http.ServeMux
}

func New() (*App, error) {
	application := &App{}

	g := generator.New()
	s := store.New()

	application.mux = http.NewServeMux()
	application.e = endpoint.New(g, s)
	application.mux.HandleFunc("/", application.e.ShortenURL)
	application.mux.HandleFunc("/{id}", application.e.Redirect)

	return application, nil
}

func (app *App) Run() error {
	log.Println("Starting server...")

	handler := middleware.Logging(app.mux)
	handler = middleware.Recovery(handler)

	if err := http.ListenAndServe(":8080", handler); err != nil {
		log.Fatal(err)
	}
	return nil
}
