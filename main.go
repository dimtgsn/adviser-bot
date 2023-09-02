package main

import (
	"flag"
	"github.com/dmitry1721/adviser-bot.git/clients/telegram"
	"log"
)

func main() {

	tgClient = telegram.New(mustHost(), mustToken())

	// fetcher = fetcher.New()

	// processor = processor.New()

	// consumer = consumer.Start(fetcher, processor)
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
