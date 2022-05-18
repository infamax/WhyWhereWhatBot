package telegram

import (
	"context"
	pb "github.com/infamax/WhyWhereWhatBot/api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"strconv"
)

func (b *Bot) getQuestionsAndAnswers(id uint64, limit int) error {
	list, err := b.client.GetQuestions(context.TODO(), &pb.Url{
		Ref: "https://db.chgk.info/xml/random/limit" + strconv.Itoa(limit),
	})
	if err != nil {
		return status.Errorf(codes.Unavailable, "not getting questions and answers")
	}
	log.Println("Answers = ", list.Answers)
	for _, answer := range list.Answers {
		log.Println("answer = ", answer)
	}
	b.cacheQuestions.SetUserQuestions(id, list.Questions, list.Answers)
	b.cacheGame.SetUserStartGame(id)
	return nil
}
