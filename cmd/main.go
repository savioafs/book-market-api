package main

import (
	"fmt"
	"github.com/savioafs/book-market/internal/config"
	"github.com/savioafs/book-market/internal/database"
	"github.com/savioafs/book-market/internal/server"
	"log"
)

func main() {
	connectionString := config.LoadConfig()

	db, err := database.InitMySQL(connectionString)
	if err != nil {
		log.Fatal("fail to init database", err)
	}

	defer db.Close()

	// TODO: return service
	//go service.MonitorCouponsByExpirationDate(db)
	fmt.Println("server is running 🚀")
	sv := server.NewServer("8080", db)

	sv.Run()
}
