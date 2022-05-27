package main

import (
	"github.com/rcrespodev/Blogs/design/repository/pkg/server"
	"log"
	"os"
)

func main() {
	srv := server.New(os.Getenv("APP_HOST"), os.Getenv("APP_PORT"))
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
