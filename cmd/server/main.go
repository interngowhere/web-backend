package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/interngowhere/web-backend/internal/middleware"
	"github.com/interngowhere/web-backend/internal/router"
)

func main() {
	s := CreateNewServer()
	r := s.Router
	middleware.Setup(r)
	router.Setup(r)
	fmt.Println("Listening on port 8000 at http://localhost:8000")

	log.Fatalln(http.ListenAndServe("127.0.0.1:8000", r))
}

type Server struct {
    Router *chi.Mux
    // Db, config can be added here
}

func CreateNewServer() *Server {
    s := &Server{}
    s.Router = chi.NewRouter()
    return s
}
