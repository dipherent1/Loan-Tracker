package Config

import (
	"context"
	"fmt"
	custommongo "loaner/CustomMongo"
	"log"
	"time"
)

func ping(client custommongo.Client, ctx context.Context) error {
	if err := client.Ping(ctx); err != nil {
		return err
	}
	fmt.Println("Connected to MongoDB successfully!")
	return nil
}

func ConnectDB() custommongo.Client {
	Envinit()
	// Connect to the database
	client, err := custommongo.NewClient(MONGO_CONNECTION_STRING)

	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	// create a context with time out
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Ping the database
	err = ping(client, ctx)
	if err != nil {
		log.Fatalf("Failed to ping MongoDB: %v", err)
	}

	return client
}
