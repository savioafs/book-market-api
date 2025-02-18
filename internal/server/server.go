package server

import (
	"github.com/gin-gonic/gin"
	"github.com/savioafs/book-market/internal/server/routes"
	"log"
)

type Server struct {
	port   string
	server *gin.Engine
}

func NewServer(port string) *Server {
	return &Server{
		port:   port,
		server: gin.Default(),
	}
}

func (s *Server) Run() {
	router := routes.ConfigRoutes(s.server)
	log.Printf("server is running on port %s", s.port)
	log.Fatal(router.Run(s.port))
}
