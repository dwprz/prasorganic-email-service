package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/dwprz/prasorganic-email-service/src/common/logger"
	"github.com/dwprz/prasorganic-email-service/src/core/broker"
	"github.com/dwprz/prasorganic-email-service/src/infrastructure/config"
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
	defer time.Sleep(5 * time.Second)

	logger := logger.New()
	appStatus := os.Getenv("PRASORGANIC_APP_STATUS")
	conf := config.New(appStatus, logger)

	gmailService := oauth.NewGmailService(conf, logger)
	emailService := service.NewEmail(gmailService, logger)

	rabbitMQClient := broker.NewRabbitMQClient(emailService, conf, logger)

	go func() {
		defer rabbitMQClient.Close()
		rabbitMQClient.Consume(ctx)
	}()

	<-ctx.Done()
}
