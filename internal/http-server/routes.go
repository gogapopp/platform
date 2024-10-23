package httpserver

import (
	"net/http"
	"platform/internal/handlers"
	"platform/ui"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func Routes() http.Handler {
	router := chi.NewRouter()

	router.Use(
		middleware.Recoverer,
		secureHeaders,
		middleware.RequestID,
		middleware.Logger,
	)

	fileServer := http.StripPrefix("/static/", http.FileServer(http.FS(ui.Files)))
	router.Get("/static/", cacheStatic(fileServer).ServeHTTP)

	router.Get("/", handlers.About)

	return router
}
