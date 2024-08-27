package main

import (
	"context"
	"fmt"
	"loaner/Config"
	"loaner/Deliverables/routers"
	"log"
)

func main() {
	// Connect to the database
	client := Config.ConnectDB()
	fmt.Println(client)
	fmt.Println(client)
	fmt.Println(client)
	fmt.Println(client)
	// Defer the closing of the database
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			log.Fatal(err)
		}
	}()

	r := routers.Setuprouter(client)
	if r != nil {
		r.Run(Config.Port)
	} else {
		log.Fatal("Failed to start server")
	}
}
