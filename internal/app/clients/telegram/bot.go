package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	pb "github.com/infamax/WhyWhereWhatBot/api"
	"github.com/infamax/WhyWhereWhatBot/internal/app/cache"
)

type Bot struct {
	bot            *tgbotapi.BotAPI
	client         pb.WhyWhereWhatServerClient
	cacheGame      *cache.UserGame
	cacheQuestions *cache.UserQuestions
}

func New(bot *tgbotapi.BotAPI, client pb.WhyWhereWhatServerClient) *Bot {
	return &Bot{
		bot:            bot,
		client:         client,
		cacheGame:      cache.NewUserGame(),
		cacheQuestions: cache.NewUserQuestions(),
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
