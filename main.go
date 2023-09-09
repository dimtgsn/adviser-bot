package main

import (
	"context"
	"flag"
	tgClient "github.com/dmitry1721/adviser-bot.git/clients/telegram"
	"github.com/dmitry1721/adviser-bot.git/consumer/event-consumer"
	"github.com/dmitry1721/adviser-bot.git/events/telegram"
	"github.com/dmitry1721/adviser-bot.git/storage/sqlite"
	"log"
)

const (
	tokenBotHost      = "api.telegram.org"
	storagePath       = "files_storage"
	sqliteStoragePath = "data/sqlite/storage.db"
	batchSize         = 100
)

func main() {

	//s := files.New(storagePath)
	s, err := sqlite.New(sqliteStoragePath)
	if err != nil {
		log.Fatal("can't connect to storage", err)
	}
	ctx := context.TODO()
	if err := s.Init(ctx); err != nil {
		log.Fatal("can't init storage", err)
	}

	eventsProcessor := telegram.New(
		tgClient.New(tokenBotHost, mustToken()),
		s,
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
