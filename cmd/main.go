package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/dwprz/prasorganic-email-service/src/core/broker"
	"github.com/dwprz/prasorganic-email-service/src/infrastructure/oauth"
	"github.com/dwprz/prasorganic-email-service/src/service"
)

func HandleCloseApp(cancel context.CancelFunc) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-c
		cancel()
	}()
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	HandleCloseApp(cancel)

	gmailService := oauth.NewGmailService()
	emailService := service.NewEmail(gmailService)

	rabbitMQClient := broker.NewRabbitMQClient(emailService)

	go func() {
		defer rabbitMQClient.Close()
		rabbitMQClient.Consume(ctx)
	}()

	<-ctx.Done()
}
