package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	jwt "github.com/interngowhere/web-backend/internal/auth"
	"github.com/interngowhere/web-backend/internal/database"
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

	// Set up server
	s := CreateNewServer()
	r := s.Router
	jwt.Init(os.Getenv("JWT_SIGKEY"))
	middleware.Setup(r)
	router.Setup(r)

	// Set up database
	database.Init()
	defer database.Client.Close()

	log.Println("Listening on port 8000 at http://localhost:8000")

	log.Fatalln(http.ListenAndServe("localhost:8000", r))
}

type Server struct {
	Router *chi.Mux
	// additional server config can be added here
}

func CreateNewServer() *Server {
	s := &Server{}
	s.Router = chi.NewRouter()
	return s
}
