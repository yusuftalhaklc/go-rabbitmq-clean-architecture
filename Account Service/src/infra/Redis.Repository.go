package infra

import (
	"context"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisRepository struct {
	redisClient *redis.Client
}

func NewRedisRepository() *RedisRepository {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		log.Fatal("[-] Error connecting to Redis:", err)
	}

	log.Println("[+] redis connection established.")

	return &RedisRepository{
		redisClient: client,
	}
}

func (r *RedisRepository) SetEmailVerificationCode(key, code string) error {
	err := r.redisClient.Set(context.Background(), key, code, 10*time.Minute).Err()
	if err != nil {
		return err
	}
	return nil
}

func (r *RedisRepository) GetEmailVerification(key string) string {
	val, err := r.redisClient.Get(context.Background(), key).Result()
	if err != nil && err != redis.Nil {
		return ""
	}
	return val
}
