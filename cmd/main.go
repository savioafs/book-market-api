package main

import (
	"github.com/savioafs/book-market/internal/config"
	"github.com/savioafs/book-market/internal/database"
	"github.com/savioafs/book-market/internal/server"
	"log"
)

func main() {
	dsn := config.LoadConfig()

	db, err := database.InitGorm(dsn)
	if err != nil {
		log.Fatal("fail to init database", err)
	}

	sv := server.NewServer(":8080", db)
	sv.Run()
}
