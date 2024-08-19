package handler

import (
	"context"

	"github.com/dwprz/prasorganic-email-service/src/common/log"
	"github.com/dwprz/prasorganic-email-service/src/service"
	"github.com/sirupsen/logrus"
)

type OtpRabbitMQ struct {
	emailService service.Email
}

func NewOtpRabbitMQ(es service.Email) *OtpRabbitMQ {
	return &OtpRabbitMQ{
		emailService: es,
	}
}

func (o *OtpRabbitMQ) ProcessMessage(ctx context.Context, msg []byte) {
	const maxRetries = 3
	for i := 0; i < maxRetries; i++ {

		if err := o.emailService.SendOtp(msg); err != nil {
			log.Logger.WithFields(logrus.Fields{"location": "handler.OtpRabbitMQ", "section": "emailService.HandleNotif"})
			continue
		}

		break
	}
}
