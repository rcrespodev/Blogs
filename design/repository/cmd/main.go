package main

import (
	"github.com/rcrespodev/Blogs/design/repository/cmd/pkg/server"
	"log"
)

const (
	host = "localhost"
	port = 8080
)

func main() {
	srv := server.New(host, port)
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
