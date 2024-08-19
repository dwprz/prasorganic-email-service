package service

import (
	"encoding/base64"
	"encoding/json"

	"github.com/dwprz/prasorganic-email-service/src/common/log"
	"github.com/dwprz/prasorganic-email-service/src/model"
	"github.com/dwprz/prasorganic-email-service/template"
	"github.com/sirupsen/logrus"
	"google.golang.org/api/gmail/v1"
)

type Email interface {
	SendOtp(data []byte) error
}

type EmailImpl struct {
	gmailService *gmail.Service
}

func NewEmail(gs *gmail.Service) Email {
	return &EmailImpl{
		gmailService: gs,
	}
}

func (s *EmailImpl) SendOtp(data []byte) error {
	otpReq := new(model.OtpRequest)

	if err := json.Unmarshal(data, otpReq); err != nil {
		log.Logger.WithFields(logrus.Fields{"location": "service.EmailImpl/SendOtp", "section": "json.Unmarshal"}).Error(err)
		return err
	}

	m := new(gmail.Message)

	tmpl := template.NewOtp(otpReq.Otp)

	emailTo := "To: " + otpReq.Email + "\r\n"
	subject := "Subject: " + "OTP Verification" + "\n"
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	msg := []byte(emailTo + subject + mime + "\n" + tmpl.String())

	m.Raw = base64.URLEncoding.EncodeToString(msg)

	if _, err := s.gmailService.Users.Messages.Send("me", m).Do(); err != nil {
		log.Logger.WithFields(logrus.Fields{"location": "service.EmailImpl/SendOtp", "section": "gmail.Send"}).Error(err)
		return err
	}

	return nil
}
