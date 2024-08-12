package config

import (
	"github.com/dwprz/prasorganic-email-service/src/common/log"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func setUpForDevelopment() *Config {
	viper := viper.New()

	viper.SetConfigFile(".env")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		log.Logger.WithFields(logrus.Fields{"location": "config.setUpForDevelopment", "section": "viper.ReadInConfig"}).Fatal(err)
	}

	oauthConf := new(oauth)
	oauthConf.ClientId = viper.GetString("GMAIL_OAUTH_CLIENT_ID")
	oauthConf.ClientSecret = viper.GetString("GMAIL_OAUTH_CLIENT_SECRET")
	oauthConf.RefreshToken = viper.GetString("GMAIL_OAUTH_REFRESH_TOKEN")

	rabbitMQConf := new(rabbitMQEmailService)
	rabbitMQConf.DSN = viper.GetString("RABBITMQ_EMAIL_SERVICE_DSN")

	return &Config{
		Oauth:    oauthConf,
		RabbitMQEmailService: rabbitMQConf,
	}
}
