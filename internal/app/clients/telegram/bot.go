package telegram

import (
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	pb "github.com/infamax/WhyWhereWhatBot/api"
	cache2 "github.com/infamax/WhyWhereWhatBot/internal/cache"
	"strconv"
)

type Bot struct {
	bot            *tgbotapi.BotAPI
	client         pb.WhyWhereWhatServerClient
	cacheGame      *cache2.UserGame
	cacheQuestions *cache2.UserQuestions
}

func New(bot *tgbotapi.BotAPI, client pb.WhyWhereWhatServerClient) *Bot {
	return &Bot{
		bot:            bot,
		client:         client,
		cacheGame:      cache2.NewUserGame(),
		cacheQuestions: cache2.NewUserQuestions(),
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

func (b *Bot) handlePosCommand(message *tgbotapi.Message) error {
	pos, err := b.client.GetPositionUser(context.TODO(), &pb.TelegramId{
		Id: uint64(message.Chat.ID),
	})
	if err != nil {
		return err
	}
	text := "Твоя позиция в рейтинге " + strconv.Itoa(int(pos.Pos))
	msg := tgbotapi.NewMessage(message.Chat.ID, text)
	b.bot.Send(msg)
	return nil
}
