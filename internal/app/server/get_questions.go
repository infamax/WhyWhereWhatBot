package server

import (
	"context"
	"encoding/xml"
	pb "github.com/infamax/WhyWhereWhatBot/api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io/ioutil"
	"net/http"
)

// Questions Структура для временого получения вопросов в формате xml
type Questions struct {
	XMLName  xml.Name `xml:"search"`
	Text     string   `xml:",chardata"`
	Question []struct {
		Text                string `xml:",chardata"`
		TourFileName        string `xml:"tourFileName"`
		TournamentFileName  string `xml:"tournamentFileName"`
		QuestionId          string `xml:"QuestionId"`
		ParentId            string `xml:"ParentId"`
		Number              string `xml:"Number"`
		Type                string `xml:"Type"`
		TypeNum             string `xml:"TypeNum"`
		TextId              string `xml:"TextId"`
		Question            string `xml:"Question"`
		Answer              string `xml:"Answer"`
		PassCriteria        string `xml:"PassCriteria"`
		Authors             string `xml:"Authors"`
		Sources             string `xml:"Sources"`
		Comments            string `xml:"Comments"`
		Rating              string `xml:"Rating"`
		RatingNumber        string `xml:"RatingNumber"`
		Complexity          string `xml:"Complexity"`
		Topic               string `xml:"Topic"`
		ProcessedBySearch   string `xml:"ProcessedBySearch"`
		ParentTextID        string `xml:"parent_text_id"`
		ParentTextId        string `xml:"ParentTextId"`
		TourId              string `xml:"tourId"`
		TournamentId        string `xml:"tournamentId"`
		TourTitle           string `xml:"tourTitle"`
		TournamentTitle     string `xml:"tournamentTitle"`
		TourType            string `xml:"tourType"`
		TournamentType      string `xml:"tournamentType"`
		TourPlayedAt        string `xml:"tourPlayedAt"`
		TournamentPlayedAt  string `xml:"tournamentPlayedAt"`
		TourPlayedAt2       string `xml:"tourPlayedAt2"`
		TournamentPlayedAt2 string `xml:"tournamentPlayedAt2"`
		Notices             string `xml:"Notices"`
	} `xml:"question"`
}

// Получаем по заданной ссылке xml файл
func getXML(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return []byte{}, status.Errorf(codes.Unavailable, "Resource unavailable")
	}

	if resp.StatusCode != http.StatusOK {
		return []byte{}, status.Errorf(codes.Code(resp.StatusCode), "Not getting resource")
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, status.Errorf(codes.Unknown, "cannot read file")
	}

	return data, nil
}

func (m MainServer) GetQuestions(ctx context.Context, req *pb.Url) (*pb.List, error) {
	xmlFile, err := getXML(req.Ref)
	if err != nil {
		return nil, status.Errorf(codes.DeadlineExceeded, "cannot to access this url")
	}
	var listQuestionsAndAnswer Questions
	err = xml.Unmarshal(xmlFile, &listQuestionsAndAnswer)
	if err != nil {
		return nil, status.Errorf(codes.Unknown, "cannot unmarshall this file")
	}

	res := new(pb.List)
	res.Questions = make([]string, len(listQuestionsAndAnswer.Question), len(listQuestionsAndAnswer.Question))
	res.Answers = make([]string, len(listQuestionsAndAnswer.Question), len(listQuestionsAndAnswer.Question))
	for i, v := range listQuestionsAndAnswer.Question {
		res.Questions[i] = v.Question
		res.Answers[i] = v.Answer
	}
	return res, nil
}
