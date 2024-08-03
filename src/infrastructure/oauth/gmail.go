package oauth

import (
	"context"
	"github.com/dwprz/prasorganic-email-service/src/infrastructure/config"
	"github.com/sirupsen/logrus"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/gmail/v1"
	"google.golang.org/api/option"
)

func NewGmailService(conf *config.Config, logger *logrus.Logger) *gmail.Service {
	ctx := context.Background()

	oauthConf := &oauth2.Config{
		ClientID:     conf.Oauth.ClientId,
		ClientSecret: conf.Oauth.ClientSecret,
		Endpoint:     google.Endpoint,
		Scopes:       []string{gmail.GmailSendScope},
	}

	token := &oauth2.Token{RefreshToken: conf.Oauth.RefreshToken}
	tokenSource := oauthConf.TokenSource(ctx, token)

	srvc, err := gmail.NewService(ctx, option.WithTokenSource(tokenSource))
	if err != nil {
		logger.WithFields(logrus.Fields{"location": "gmail.NewService", "section": "gmail.NewService"}).Error(err)
	}

	return srvc
}
