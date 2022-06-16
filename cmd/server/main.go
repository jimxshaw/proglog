package main

import (
	"log"

	"github.com/jimxshaw/proglog/internal/server"
)

func main() {
	server := server.NewHttpServer(":8000")
	log.Fatal(server.ListenAndServe())
}
