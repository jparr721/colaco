package main

import (
	"colaco-server/internal/db"
	"colaco-server/internal/routes"
	"context"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	logger := zap.Must(zap.NewDevelopment())
	if os.Getenv("APP_ENV") == "production" {
		logger = zap.Must(zap.NewProduction())
	}
	zap.ReplaceGlobals(logger)
}

func dbContext(db *db.ColacoDB) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(r.Context(), "db", db)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func main() {
	database := &db.ColacoDB{}
	database.Init()

	r := chi.NewRouter()

	r.Use(dbContext(database))
	r.Use(middleware.Logger)
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	r.Route("/v1", routes.MakeColaCoRouter)

	zap.L().Info("Server running on port 8000")
	http.ListenAndServe(":8000", r)
}
