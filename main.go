package main

import (
	"github.com/sdq-codes/maze-app/server"
	"log"
)

func main() {
	srv := server.Server()

	log.Println("Server listening on", srv.Addr)
	log.Fatal(srv.ListenAndServe())
}
