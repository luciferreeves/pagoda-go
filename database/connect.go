package database

import (
	"log"
	"os"
	"pagoda/models"

	"github.com/redis/go-redis/v9"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() {
	var err error // define error here to prevent overshadowing the global DB

	env := os.Getenv("DSN")
	DB, err = gorm.Open(postgres.Open(env), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	err = DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal(err)
	}
}

func ConnectRedis() {
	Cache = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDR"),
		Password: os.Getenv("REDIS_PASSWORD"),
	})

	_, err := Cache.Ping(Ctx).Result()

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Redis connected")
}
