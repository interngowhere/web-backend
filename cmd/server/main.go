package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	jwt "github.com/interngowhere/web-backend/internal/auth"
	"github.com/interngowhere/web-backend/internal/database"
	"github.com/interngowhere/web-backend/internal/middleware"
	"github.com/interngowhere/web-backend/internal/router"
	"github.com/interngowhere/web-backend/internal/server"

	_ "github.com/lib/pq"
)

func main() {
	// Set up server
	s := server.New()
	r := s.Router
	jwt.Init(os.Getenv("JWT_SIGKEY"))
	middleware.Setup(r)
	router.Setup(r)

	// Set up database
	database.Init()
	defer database.Client.Close()

	log.Println(fmt.Printf("Listening on port 8000 at http://%s:8000\n", os.Getenv("SERVER_HOST")))

	log.Fatalln(http.ListenAndServe(fmt.Sprintf("%s:8000", os.Getenv("SERVER_HOST")), r))
}

