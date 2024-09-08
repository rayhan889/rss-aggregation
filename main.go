package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	"github.com/rayhan889/rss-aggr/handle_error"
	"github.com/rayhan889/rss-aggr/internal/database"
	"github.com/rayhan889/rss-aggr/readiness"
	"github.com/rayhan889/rss-aggr/users"

	_ "github.com/lib/pq"
)

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

	apiCfg := &users.ApiConfig{
		DB: database.New(conn),
	}

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
	v1Router.Post("/users", apiCfg.HandleCreateNewUser)

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