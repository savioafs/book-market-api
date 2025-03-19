package server

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/savioafs/book-market/internal/server/routes"

	"log"
)

type Server struct {
	port   string
	server *gin.Engine
	DB     *sql.DB
}

func NewServer(port string, db *sql.DB) *Server {

	return &Server{
		port:   port,
		server: gin.Default(),
		DB:     db,
	}
}

func (s *Server) Run() {
	routes.SetupRoutes(s.server, s.DB)
	log.Printf("server is running on port %s", s.port)
	if err := s.server.Run(":" + s.port); err != nil {
		log.Fatalf("failed to start server %v", err)
	}

}
