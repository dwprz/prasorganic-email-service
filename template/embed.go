package template

import (
	_ "embed"
	"html/template"
	"strings"
	"github.com/sirupsen/logrus"
)

//go:embed html/otp.html
var otpEmbed string

func NewOtp(logger *logrus.Logger, otp string) *strings.Builder {
	t, err := template.New("otp").Parse(otpEmbed)
	if err != nil {
		logger.WithFields(logrus.Fields{"location": "template.NewOtp", "section": "template.Parse"}).Error(err)
	}

	var body strings.Builder

	if err := t.Execute(&body, map[string]string{"otp": otp}); err != nil {
		logger.WithFields(logrus.Fields{"location": "template.NewOtp", "section": "template.Execute"}).Error(err)
	}

	return &body
}
