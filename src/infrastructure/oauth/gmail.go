package oauth

import (
	"context"

	"github.com/dwprz/prasorganic-email-service/src/common/log"
	"github.com/dwprz/prasorganic-email-service/src/infrastructure/config"
	"github.com/sirupsen/logrus"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/gmail/v1"
	"google.golang.org/api/option"
)

func NewGmailService() *gmail.Service {
	ctx := context.Background()

	oauthConf := &oauth2.Config{
		ClientID:     config.Conf.Oauth.ClientId,
		ClientSecret: config.Conf.Oauth.ClientSecret,
		Endpoint:     google.Endpoint,
		Scopes:       []string{gmail.GmailSendScope},
	}

	token := &oauth2.Token{RefreshToken: config.Conf.Oauth.RefreshToken}
	tokenSource := oauthConf.TokenSource(ctx, token)

	srvc, err := gmail.NewService(ctx, option.WithTokenSource(tokenSource))
	if err != nil {
		log.Logger.WithFields(logrus.Fields{"location": "oauth.NewGmailService", "section": "gmail.NewService"}).Error(err)
	}

	return srvc
}
