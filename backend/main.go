package main

import (
	"fmt"
	"log"
	"os"

	"personal-site/server"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("No arguments given. Please pass the path to the client code.")
	}

	pathToStaticWebsite := os.Args[1]

	fmt.Println("Starting Server")
	server.Run(pathToStaticWebsite)
}
