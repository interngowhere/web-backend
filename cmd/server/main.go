package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/interngowhere/web-backend/ent"
	"github.com/interngowhere/web-backend/internal/middleware"
	"github.com/interngowhere/web-backend/internal/router"
	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)
  
func main() {
	// load .env file
	err := godotenv.Load(".env")
  
	if err != nil {
	  log.Fatalf("Failed to load .env file: %v", err)
	}
	
	s := CreateNewServer()
	r := s.Router
	middleware.Setup(r)
	router.Setup(r)

    defer s.DbClient.Close()

	fmt.Println("Listening on port 8000 at http://localhost:8000")

	log.Fatalln(http.ListenAndServe("localhost:8000", r))
}

type Server struct {
    Router *chi.Mux
	DbClient *ent.Client
    // additional server config can be added here
}

func CreateNewServer() *Server {
    s := &Server{}
    s.Router = chi.NewRouter()

	// Setup database
	err := error(nil)
	s.DbClient, err = ent.Open(
		os.Getenv("DB_DIALECT"), 
		fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", 
			os.Getenv("DB_HOST"), 
			os.Getenv("DB_PORT"),
			os.Getenv("DB_USER"),
			os.Getenv("DB_NAME"),
			os.Getenv("DB_PASSWORD")))
    
	if err != nil {
        log.Fatalf("Failed to establish connection to database: %v", err)
    }

	// Run the auto migration tool.
    if err := s.DbClient.Schema.Create(context.Background()); err != nil {
        log.Fatalf("Failed to create schema resources: %v", err)
    }

    return s
}
