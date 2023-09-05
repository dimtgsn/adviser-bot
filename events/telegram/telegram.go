package telegram

import "github.com/dmitry1721/adviser-bot.git/clients/telegram"

type Processor struct {
	tg     *telegram.Client
	offset int
	// storage
}

func New(client *telegram.Client) Processor {
	return Processor{}
}
