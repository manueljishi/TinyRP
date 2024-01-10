package main

import (
	"fmt"
	"log"

	"github.com/manueljishi/tinyrp/internal/server"
)

func main() {
	fmt.Println("hola")
	if err := server.Run(); err != nil {
		log.Fatalf("could not start the server %v", err)
	}
	fmt.Println("adios")
}
