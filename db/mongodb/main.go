package mongodb

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connect() (*mongo.Client, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("erro loading .env file: %v", err)
	}

	mongoURL := os.Getenv("MONGO_URL")
	dbName := os.Getenv("MONGO_DB")
	username := os.Getenv("MONGO_USER")
	password := os.Getenv("MONGO_PASS")

	client, err := connectToMongoDB(mongoURL, dbName, username, password)
	if err != nil {
		return nil, fmt.Errorf("error connecting to MongoDB: %v", err)
	}

	log.Println("Connected to MongoDB")

	return client, nil
}

func connectToMongoDB(url, dbName, username, password string) (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI(url).SetAuth(options.Credential{
		Username: username,
		Password: password,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, fmt.Errorf("error connecting to MongoDB server: %v", err)
	}

	return client, nil
}
