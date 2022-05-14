package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"strconv"
)

const (
	commandStart = "start"
	commandHelp  = "help"
	commandTop   = "top"
	commandRules = "rules"
	commandMyPos = "mypos"
	commandGame  = "game"
)

func (b *Bot) handleUpdates(updates tgbotapi.UpdatesChannel) {
	for update := range updates {
		if update.Message == nil {
			continue
		}
		if update.Message.IsCommand() {
			b.handleCommand(update.Message)
			continue
		} else if b.cacheGame.IsUserPlayGame(uint64(update.Message.Chat.ID)) {
			b.checkAnswer(update.Message, b.cacheQuestions.GetCountAskedQuestions(uint64(update.Message.Chat.ID)))
		} else {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID,
				"Введите комманду, чтобы начать игру! Узнать список моих комманд /help")
			_, err := b.bot.Send(msg)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}

func (b *Bot) handleCommand(message *tgbotapi.Message) error {
	if b.cacheGame.IsUserPlayGame(uint64(message.Chat.ID)) {
		msg := tgbotapi.NewMessage(message.Chat.ID, "Ты ввел комманду. Поэтому я вынужден прервать начатую тобой игру.")
		b.cacheGame.SetUserStopGame(uint64(message.Chat.ID))
		_, err := b.bot.Send(msg)
		return err
	}
	msg := tgbotapi.NewMessage(message.Chat.ID, "")
	switch message.Command() {
	case commandStart:
		return b.handleStartCommand(message, "Привет! Чтобы узнать мои возможности введи комманду /help")
	case commandHelp:
		return b.handleTemplateCommand(message, Help)
	case commandRules:
		return b.handleTemplateCommand(message, Rules)
	case commandTop:
		return b.handleTopCommand(message, 3)
	case commandMyPos:
		return b.handlePosCommand(message)
	case commandGame:
		err := b.getQuestionsAndAnswers(uint64(message.Chat.ID), 10)
		if err != nil {
			text := "К сожалению я не могу получить список вопросов с сервиса. Попробуй воспользоваться ботом позже"
			msg := tgbotapi.NewMessage(message.Chat.ID, text)
			b.bot.Send(msg)
			return err
		}
		return b.playGame(message, 0)
	default:
		msg.Text = "Ты ввел неизвестную комманду"
		_, err := b.bot.Send(msg)
		return err
	}
}

func (b *Bot) handleStartCommand(message *tgbotapi.Message, greeting string) error {
	_ = b.initUser(message.Chat.ID, message.Chat.UserName)
	msg := tgbotapi.NewMessage(message.Chat.ID, greeting)
	_, err := b.bot.Send(msg)
	return err
}

func (b *Bot) handleTopCommand(message *tgbotapi.Message, limit int) error {
	// TODO нехватка игровок обеспечить проверку
	users, err := b.getTop(limit)
	text := "Список лидеров\n"
	for i, user := range users.Name {
		text += strconv.Itoa(i+1) + ". @" + user + " Score: " + strconv.Itoa(int(users.Score[i].Count)) + "\n"
	}
	msg := tgbotapi.NewMessage(message.Chat.ID, text)
	_, err = b.bot.Send(msg)
	return err
}

func (b *Bot) handleTemplateCommand(message *tgbotapi.Message, text string) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, text)
	_, err := b.bot.Send(msg)
	return err
}

func (b *Bot) handleMessage(message *tgbotapi.Message) {
	log.Printf("[%s] %s", message.From.UserName, message.Text)

	msg := tgbotapi.NewMessage(message.Chat.ID, message.Text)
	//msg.ReplyToMessageID = message.MessageID
	b.bot.Send(msg)
}
