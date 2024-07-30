package account

type AccountRepository interface {
	Create(account Account) error
	Update(email string, verified bool) error
}

type RedisRepository interface {
	SetEmailVerificationCode(key, code string) error
	GetEmailVerification(key string) string
}

type RabbitMQProducer interface {
	Publish(body []byte) error
}

type Services struct {
	AccountRepo      AccountRepository
	RedisRepo        RedisRepository
	RabbitMqProducer RabbitMQProducer
}
