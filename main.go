package main

import (
	"flag"
	tgClient "github.com/dmitry1721/adviser-bot.git/clients/telegram"
	"github.com/dmitry1721/adviser-bot.git/consumer/event-consumer"
	"github.com/dmitry1721/adviser-bot.git/events/telegram"
	"github.com/dmitry1721/adviser-bot.git/storage/files"
	"log"
)

const (
	storagePath = "storage"
	batchSize   = 100
)

func main() {

	eventsProcessor := telegram.New(
		tgClient.New(mustHost(), mustToken()),
		files.New(storagePath),
	)

	log.Print("service started")

	consumer := event_consumer.New(eventsProcessor, eventsProcessor, batchSize)
	if err := consumer.Start(); err != nil {
		log.Fatal("service stopped", err)
	}
}

func mustToken() string {
	token := flag.String(
		"tg-bot-token",
		"",
		"token for access to telegram bot",
	)
	flag.Parse()

	if *token == "" {
		log.Fatal("token id not specified")
	}

	return *token
}

func mustHost() string {
	host := flag.String(
		"tg-bot-host",
		"",
		"telegram bot host",
	)
	flag.Parse()

	if *host == "" {
		log.Fatal("host not specified")
	}

	return *host
}
