package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	routes "github.com/Dimoonevs/SportsApp/logger-service/internal/api"
	"github.com/Dimoonevs/SportsApp/logger-service/internal/repository"
	"github.com/Dimoonevs/SportsApp/logger-service/internal/service"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	webPort  = "8001"
	mongoURL = "mongodb://mongo:27017"
)

func main() {

	// connect to mongo
	mongoClient, err := connectToMongo(mongoURL)
	if err != nil {
		log.Println("Error connecting to MongoDB!")
		log.Panic(err)
	}
	log.Println("Connected to MongoDB!")

	// create a context in order to disconnect
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	defer func() {
		if err = mongoClient.Disconnect(ctx); err != nil {
			log.Println("Error disconnecting from MongoDB!")
			panic(err)
		}
	}()

	// create service and repository
	var repo repository.RepositoryLogs
	repo = &repository.RepositoryLogsMongo{
		Mongo: mongoClient,
	}

	service := &service.LogsService{
		Repo: repo,
	}

	app := routes.AppLogger{
		Service: service,
	}

	log.Printf("Starting Logger Service on port %s", webPort)
	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.Routes(),
	}
	err = server.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}

}

func connectToMongo(urlMongo string) (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI(urlMongo)

	clientOptions.SetAuth(options.Credential{
		Username: "admin",
		Password: "password",
	})

	c, err := mongo.Connect(context.Background(), clientOptions)

	if err != nil {
		return nil, err
	}
	return c, nil

}
