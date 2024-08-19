package consumer

import (
	"context"

	"github.com/dwprz/prasorganic-email-service/src/common/log"
	"github.com/dwprz/prasorganic-email-service/src/core/broker/handler"
	"github.com/dwprz/prasorganic-email-service/src/infrastructure/config"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/sirupsen/logrus"
)

type OtpRabbitMQ struct {
	otpHandler *handler.OtpRabbitMQ
	connection *amqp.Connection
}

func NewOtpRabbitMQ(oh *handler.OtpRabbitMQ) *OtpRabbitMQ {
	conn, err := amqp.Dial(config.Conf.RabbitMQEmailService.DSN)
	if err != nil {
		log.Logger.WithFields(logrus.Fields{"location": "broker.OtpRabbitMQ/Consume", "section": "amqp.Dial"}).Fatal(err)
	}

	return &OtpRabbitMQ{
		otpHandler: oh,
		connection: conn,
	}
}

func (o *OtpRabbitMQ) Consume(ctx context.Context) {
	log.Logger.Info("rabbitmq client start consume")

	channel, err := o.connection.Channel()
	if err != nil {
		log.Logger.WithFields(logrus.Fields{"location": "consumer.OtpRabbitMQ/Consume", "section": "conn.Channel"}).Fatal(err)
	}

	defer channel.Close()

	otpConsumer, err := channel.ConsumeWithContext(ctx, "otp", "otp-consumer", true, false, false, false, nil)
	if err != nil {
		log.Logger.WithFields(logrus.Fields{"location": "consumer.OtpRabbitMQ/Consume", "section": "channel.ConsumeWithContext"}).Fatal(err)
	}

	for {
		select {
		case msg := <-otpConsumer:

			go func(msg amqp.Delivery) {
				defer func() {
					if o := recover(); o != nil {
						log.Logger.WithFields(logrus.Fields{"location": "consumer.OtpRabbitMQ/Consume", "section": "ProcessMessage"}).Errorf("Recovered from panic: %v", o)
					}
				}()

				o.otpHandler.ProcessMessage(ctx, msg.Body)
			}(msg)

		case <-ctx.Done():
			return
		}
	}
}

func (o *OtpRabbitMQ) Close() {
	if err := o.connection.Close(); err != nil {
		log.Logger.WithFields(logrus.Fields{"location": "consumer.OtpRabbitMQ/Close", "section": "connection.Close"}).Error(err)
	}
}
