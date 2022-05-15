package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	pb "github.com/infamax/WhyWhereWhatBot/api"
	cache2 "github.com/infamax/WhyWhereWhatBot/internal/cache"
)

type Bot struct {
	bot            *tgbotapi.BotAPI
	client         pb.WhyWhereWhatServerClient
	cacheGame      *cache2.UserGame
	cacheQuestions *cache2.UserQuestions
	cacheTime      *cache2.TimeUsers
}

func New(bot *tgbotapi.BotAPI, client pb.WhyWhereWhatServerClient) *Bot {
	return &Bot{
		bot:            bot,
		client:         client,
		cacheGame:      cache2.NewUserGame(),
		cacheQuestions: cache2.NewUserQuestions(),
		cacheTime:      cache2.NewUserTimer(),
	}
}

func (b *Bot) Start() error {
	updates := b.initUpdatesChannel()
	b.handleUpdates(updates)
	return nil
}

func (b *Bot) initUpdatesChannel() tgbotapi.UpdatesChannel {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	return b.bot.GetUpdatesChan(u)
}
