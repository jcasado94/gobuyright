package server

import (
	"gobuyright/pkg/entity"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// Server controls the routing of the incoming queries.
type Server struct {
	router *mux.Router
}

// NewServer creates a new server given the needed Services.
func NewServer(iuserService entity.IUserService, usService entity.UsageSelectionService) *Server {
	s := Server{router: mux.NewRouter()}
	NewIUserRouter(iuserService, s.newSubrouter("/iuser"))
	NewUsageSelectionRouter(usService, s.newSubrouter("/usageselection"))
	return &s
}

// Start starts the routing to port 8080.
func (s *Server) Start() {
	log.Println("Listening on port 8080")
	if err := http.ListenAndServe(":8080", handlers.LoggingHandler(os.Stdout, s.router)); err != nil {
		log.Fatal("http.ListenAndServe: ", err)
	}
}

func (s *Server) newSubrouter(path string) *mux.Router {
	return s.router.PathPrefix(path).Subrouter()
}
