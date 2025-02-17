package main

import "github.com/savioafs/book-market/server"

func main() {

	sv := server.NewServer(":8080")
	sv.Run()
}
