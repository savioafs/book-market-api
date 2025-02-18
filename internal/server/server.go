package server

import (
	"github.com/gin-gonic/gin"
	"github.com/savioafs/book-market/internal/server/routes"
	"gorm.io/gorm"
	"log"
)

type Server struct {
	port   string
	server *gin.Engine
	DB     *gorm.DB
}

func NewServer(port string, db *gorm.DB) *Server {

	return &Server{
		port:   port,
		server: gin.Default(),
		DB:     db,
	}
}

func (s *Server) Run() {
	router := routes.ConfigRoutes(s.server, s.DB)
	log.Printf("server is running on port %s", s.port)
	log.Fatal(router.Run(s.port))
}
