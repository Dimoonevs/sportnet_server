package repository

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/jackc/pgx/v4/pgxpool"
)

var counts int64

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

func ConnectToRedis() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})

	pong, err := client.Ping(context.Background()).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to Redis:", pong)

	return client
}
