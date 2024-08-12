package config

import (
	"context"
	"os"
	"strings"

	"github.com/dwprz/prasorganic-email-service/src/common/log"
	vault "github.com/hashicorp/vault/api"
	"github.com/sirupsen/logrus"
)

func setUpForNonDevelopment(appStatus string) *Config {
	vaultConf := vault.DefaultConfig()
	vaultConf.Address = os.Getenv("PRASORGANIC_CONFIG_ADDRESS")

	client, err := vault.NewClient(vaultConf)
	if err != nil {
		log.Logger.Fatalf("vault new client: %v", err)
	}

	client.SetToken(os.Getenv("PRASORGANIC_CONFIG_TOKEN"))
	mountPath := "prasorganic-secrets" + "-" + strings.ToLower(appStatus)

	oauthSecrets, err := client.KVv2(mountPath).Get(context.Background(), "oauth")
	if err != nil {
		log.Logger.WithFields(logrus.Fields{"location": "config.setUpForNonDevelopment", "section": "KVv2.Get"}).Fatal(err)
	}

	rabbitMQEmailServiceSecrets, err := client.KVv2(mountPath).Get(context.Background(), "rabbitmq-email-service")
	if err != nil {
		log.Logger.WithFields(logrus.Fields{"location": "config.setUpForNonDevelopment", "section": "KVv2.Get"}).Fatal(err)
	}

	oauthConf := new(oauth)
	oauthConf.ClientId = oauthSecrets.Data["GMAIL_CLIENT_ID"].(string)
	oauthConf.ClientSecret = oauthSecrets.Data["GMAIL_CLIENT_SECRET"].(string)
	oauthConf.RefreshToken = oauthSecrets.Data["GMAIL_REFRESH_TOKEN"].(string)

	rabbitMQEmailServiceConf := new(rabbitMQEmailService)
	rabbitMQEmailServiceConf.DSN = rabbitMQEmailServiceSecrets.Data["DSN"].(string)

	return &Config{
		Oauth:                oauthConf,
		RabbitMQEmailService: rabbitMQEmailServiceConf,
	}
}
