package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type APIConfig struct {
	DB *sql.DB
}

func main() {
	godotenv.Load()
	var PORT string = os.Getenv("PORT")

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DB_URL not set")
	}

	conn, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal(err)
	}
	apiConfig := &APIConfig{
		DB: conn,
	}

    go startScrap(apiConfig, 10,time.Minute)

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))
    router.Use(middleware.Logger)

	v1Router := chi.NewRouter()
	v1Router.Get("/ready", handlerReadiness)
	v1Router.Get("/err", handleError)
	v1Router.Post("/users", apiConfig.handleCreateUser)
	v1Router.Get("/users", apiConfig.middlewareAuth(apiConfig.handleGetUser))
    v1Router.Post("/feeds", apiConfig.middlewareAuth(apiConfig.handleCreateFeed))
    v1Router.Get("/feeds", apiConfig.handleGetFeed)
    v1Router.Post("/feed_follows", apiConfig.middlewareAuth(apiConfig.handleCreateFeedFollow))
    v1Router.Get("/feed_follows", apiConfig.middlewareAuth(apiConfig.handleGetFeedFollow))
    v1Router.Delete("/feed_follows/{feedid}", apiConfig.middlewareAuth(apiConfig.handleUnFollowFeed))

	router.Mount("/v1", v1Router)

	server := &http.Server{
		Handler: router,
		Addr:    ":" + PORT,
	}
	log.Printf("Serving on port: %s\n", PORT)
	log.Fatal(server.ListenAndServe())
}
