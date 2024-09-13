package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	"github.com/rayhan889/rss-aggr/feeds"
	"github.com/rayhan889/rss-aggr/handle_error"
	"github.com/rayhan889/rss-aggr/internal/database"
	"github.com/rayhan889/rss-aggr/readiness"
	"github.com/rayhan889/rss-aggr/users"

	_ "github.com/lib/pq"
)

type ApiConfig struct {
	DB *database.Queries
}

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	portURL := os.Getenv("PORT")

	if portURL == "" {
		log.Fatal("PORT environment variable is not set")
	}

	dbURL := os.Getenv("DB_URL")

	if dbURL == "" {
		log.Fatal("DB_URL environment variable is not set")
	}

	conn, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("Can't connect to database: ", err)
	}

	apiConfig := &ApiConfig{
		DB: database.New(conn),
	}

	userConfig := &users.ApiConfig{DB: apiConfig.DB}
	feedConfig := &feeds.ApiConfig{DB: apiConfig.DB}

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	v1Router := chi.NewRouter()
	v1Router.Get("/healthz", readiness.HandleReadiness)
	v1Router.Get("/err", handle_error.HandleError)

	v1Router.Post("/users", userConfig.HandleCreateNewUser)
	v1Router.Get("/users", apiConfig.authMiddleware(userConfig.HandleGetUserByAPIKey))

	v1Router.Get("/feeds", feedConfig.HandleGetFeeds)
	v1Router.Post("/feeds", apiConfig.authMiddleware(feedConfig.HandleCreateNewFeed))
	v1Router.Get("/feeds/user/{userID}", apiConfig.authMiddleware(feedConfig.HandleGetFeedsByUserID))

	router.Mount("/v1", v1Router)

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + portURL,
	}

	log.Printf("Server started on port %s", portURL)
	errSrv := srv.ListenAndServe()

	if(errSrv != nil) {
		log.Fatal(errSrv)
	}
}