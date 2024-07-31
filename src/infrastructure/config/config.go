package config

import (
	"github.com/sirupsen/logrus"
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

func New(appStatus string, logger *logrus.Logger) *Config {
	var config *Config

	if appStatus == "DEVELOPMENT" {

		config = setUpForDevelopment(logger)
		return config
	}

	config = setUpForNonDevelopment(appStatus, logger)
	return config
}
