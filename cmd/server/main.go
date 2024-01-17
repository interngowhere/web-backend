package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	jwt "github.com/interngowhere/web-backend/internal/auth"
	"github.com/interngowhere/web-backend/internal/database"
	"github.com/interngowhere/web-backend/internal/middleware"
	"github.com/interngowhere/web-backend/internal/router"

	_ "github.com/lib/pq"
)

func main() {
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

	log.Fatalln(http.ListenAndServe(fmt.Sprintf("%s:8000", os.Getenv("SERVER_HOST")), r))
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
