package config

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type AppConfig struct {
	DB   *mongo.Database
	Port string
}

func Load() *AppConfig {
	_ = godotenv.Load()

	mongoURI := os.Getenv("MONGO_URI")
	dbName := os.Getenv("MONGO_DB")
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatal("mongo connect error:", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := client.Ping(ctx, nil); err != nil {
		log.Fatal("mongo ping error:", err)
	}

	return &AppConfig{
		DB:   client.Database(dbName),
		Port: port,
	}
}
