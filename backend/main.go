package main

import (
	server "github.com/kushian01100111/Medical-ER/Backend/package"
)

func main() {
	srv := server.NewServer(":8080")
	srv.Run()
}
