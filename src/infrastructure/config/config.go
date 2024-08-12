package config

import (
	"os"
)

type oauth struct {
	ClientId     string
	ClientSecret string
	RefreshToken string
}

type rabbitMQEmailService struct {
	DSN string
}

type Config struct {
	Oauth                *oauth
	RabbitMQEmailService *rabbitMQEmailService
}

var Conf *Config

// *config ini hanya berisi env variable
func init() {
	appStatus := os.Getenv("PRASORGANIC_APP_STATUS")

	if appStatus == "DEVELOPMENT" {

		Conf = setUpForDevelopment()
		return
	}

	Conf = setUpForNonDevelopment(appStatus)
}
