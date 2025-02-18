package main

import (
	"github.com/savioafs/book-market/internal/server"
)

func main() {

	sv := server.NewServer(":8080")
	sv.Run()
}
