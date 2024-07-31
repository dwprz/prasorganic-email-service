package service

import (
	"encoding/base64"
	"encoding/json"
	"github.com/dwprz/prasorganic-email-service/src/model"
	"github.com/dwprz/prasorganic-email-service/template"
	"github.com/sirupsen/logrus"
	"google.golang.org/api/gmail/v1"
)

type Email interface {
	SendOtp(data []byte)
}

type EmailImpl struct {
	gmailService *gmail.Service
	logger       *logrus.Logger
}

func NewEmail(gs *gmail.Service, l *logrus.Logger) Email {
	return &EmailImpl{
		gmailService: gs,
		logger:       l,
	}
}

func (s *EmailImpl) SendOtp(data []byte) {
	otpReq := new(model.OtpRequest)

	if err := json.Unmarshal(data, otpReq); err != nil {
		s.logger.WithFields(logrus.Fields{"location": "service.EmailImpl/SendOtp", "section": "json.Unmarshal"}).Error(err)
		return
	}

	m := new(gmail.Message)

	tmpl := template.NewOtp(s.logger, otpReq.Otp)

	emailTo := "To: " + otpReq.Email + "\r\n"
	subject := "Subject: " + "OTP Verification" + "\n"
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	msg := []byte(emailTo + subject + mime + "\n" + tmpl.String())

	m.Raw = base64.URLEncoding.EncodeToString(msg)

	if _, err := s.gmailService.Users.Messages.Send("me", m).Do(); err != nil {
		s.logger.WithFields(logrus.Fields{"location": "service.EmailImpl/SendOtp", "section": "gmail.Send"}).Error(err)
	}
}
