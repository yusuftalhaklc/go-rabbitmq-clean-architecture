package main

import (
	"account"
	"infra"

	"github.com/gin-gonic/gin"
)

var (
	usecase account.UseCase
)

func main() {
	app := gin.Default()

	usecase = account.NewInteractor(
		account.Services{
			AccountRepo:      infra.NewAccountRepository(),
			RedisRepo:        infra.NewRedisRepository(),
			RabbitMqProducer: infra.NewRabbitMQProducer(),
		},
	)

	app.POST("/account", Register)
	app.GET("/account", Verify)

	app.Run("localhost:3000")
}
