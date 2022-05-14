package telegram

import (
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	pb "github.com/infamax/WhyWhereWhatBot/api"
	"strconv"
)

func (b *Bot) playGame(message *tgbotapi.Message, num int) error {
	if num == 0 {
		text := "Начнем игру"
		msg := tgbotapi.NewMessage(message.Chat.ID, text)
		b.bot.Send(msg)
		b.cacheGame.SetUserStartGame(uint64(message.Chat.ID))
	}
	b.askQuestion(message, num)
	return nil
}

func (b *Bot) askQuestion(message *tgbotapi.Message, num int) {
	text := strconv.Itoa(num+1) + " вопрос " + b.cacheQuestions.GetUserQuestion(uint64(message.Chat.ID), num)
	msg := tgbotapi.NewMessage(message.Chat.ID, text)
	b.bot.Send(msg)
}

func (b *Bot) checkAnswer(message *tgbotapi.Message, num int) bool {
	var text string
	res := true
	if message.Text == b.cacheQuestions.GetUserQuestion(uint64(message.Chat.ID), num) {
		text = "Твой ответ правильный. Еще один балл в копилку"
		b.cacheQuestions.IncUserCorrectAns(uint64(message.Chat.ID))
	} else {
		text = "Твой ответ к сожалению неверный :("
		res = false
	}
	msg := tgbotapi.NewMessage(message.Chat.ID, text)
	b.bot.Send(msg)
	b.cacheQuestions.IncUserAskedQuestions(uint64(message.Chat.ID))
	if b.cacheQuestions.IsEndUserQuestions(uint64(message.Chat.ID)) {
		text := "Конец игры. Твой результат " +
			strconv.Itoa(b.cacheQuestions.GetCorrectAnsUser(uint64(message.Chat.ID))) + " баллов"
		count := b.cacheQuestions.GetCorrectAnsUser(uint64(message.Chat.ID))
		msg := tgbotapi.NewMessage(message.Chat.ID, text)
		b.bot.Send(msg)
		b.cacheGame.SetUserStopGame(uint64(message.Chat.ID))
		b.cacheQuestions.DeleteUser(uint64(message.Chat.ID))
		b.client.UpdateUserScore(context.TODO(), &pb.UserTelegram{
			TelegramId: uint64(message.Chat.ID),
			NewScore:   uint64(count),
		})
	} else {
		b.playGame(message, b.cacheQuestions.GetCountAskedQuestions(uint64(message.Chat.ID)))
	}
	return res
}
