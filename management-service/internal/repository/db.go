package repository

import (
	"context"
	"database/sql"
	"log"
	"os"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var counts int64

const (
	mongoURL = "mongodb://mongo:27017"
)

func ConnectToMongo() *mongo.Client {
	mongoClient, err := openMongo(mongoURL)
	if err != nil {
		log.Println("Error connecting to MongoDB!")
		log.Panic(err)
	}
	log.Println("Connected to MongoDB!")
	return mongoClient
}

func openMongo(urlMongo string) (*mongo.Client, error) {
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

func ConnectToPostgres() *pgxpool.Pool {
	dsn := os.Getenv("DSN")

	for {
		config, err := pgxpool.ParseConfig(dsn)
		if err != nil {
			log.Println("Postgres not yet ready...")
			counts++
		}
		pool, err := pgxpool.ConnectConfig(context.Background(), config)
		if err != nil {
			log.Println("Postgres not yet ready...")
			counts++
		} else {
			log.Println("Connected to Postgres!")
			return pool
		}

		if counts > 10 {
			log.Println(err)
			return nil
		}

		log.Println("Backing off for two seconds...")
		time.Sleep(2 * time.Second)
		continue
	}
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
