package broker

import (
	"context"

	"github.com/dwprz/prasorganic-email-service/src/common/log"
	"github.com/dwprz/prasorganic-email-service/src/infrastructure/config"
	"github.com/dwprz/prasorganic-email-service/src/service"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/sirupsen/logrus"
)

type RabbitMQClient struct {
	emailService service.Email
	connection   *amqp.Connection
}

func NewRabbitMQClient(es service.Email) *RabbitMQClient {
	conn, err := amqp.Dial(config.Conf.RabbitMQEmailService.DSN)
	if err != nil {
		log.Logger.WithFields(logrus.Fields{"location": "broker.RabbitMQClient/Consume", "section": "amqp.Dial"}).Fatal(err)
	}

	return &RabbitMQClient{
		emailService: es,
		connection:   conn,
	}
}

func (r *RabbitMQClient) Consume(ctx context.Context) {
	log.Logger.Info("rabbitmq client start consume")

	channel, err := r.connection.Channel()
	if err != nil {
		log.Logger.WithFields(logrus.Fields{"location": "broker.RabbitMQClient/Consume", "section": "conn.Channel"}).Fatal(err)
	}

	defer channel.Close()

	otpConsumer, err := channel.ConsumeWithContext(ctx, "otp", "otp-consumer", true, false, false, false, nil)
	if err != nil {
		log.Logger.WithFields(logrus.Fields{"location": "broker.RabbitMQClient/Consume", "section": "channel.ConsumeWithContext"}).Fatal(err)
	}

	for {
		select {
		case message := <-otpConsumer:
			r.emailService.SendOtp(message.Body)
		case <-ctx.Done():
			return
		}
	}
}

func (r *RabbitMQClient) Close() {
	if err := r.connection.Close(); err != nil {
		log.Logger.WithFields(logrus.Fields{"location": "broker.RabbitMQClient/Close", "section": "connection.Close"}).Error(err)
	}
}
