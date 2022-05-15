package cache

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"sync"
)

type UserQuestions struct {
	mu              sync.RWMutex
	questions       map[uint64][]string
	answers         map[uint64][]string
	correctAnsCount map[uint64]int
	numQuestions    map[uint64]int
}

func NewUserQuestions() *UserQuestions {
	return &UserQuestions{
		questions:       make(map[uint64][]string),
		answers:         make(map[uint64][]string),
		correctAnsCount: make(map[uint64]int),
		numQuestions:    make(map[uint64]int),
	}
}

func (u *UserQuestions) SetUserQuestions(id uint64, questions, answers []string) {
	defer u.mu.Unlock()
	u.mu.Lock()
	u.numQuestions[id] = 0
	u.questions[id] = questions
	u.answers[id] = answers
	u.correctAnsCount[id] = 0
}

func (u *UserQuestions) GetUserQuestion(id uint64, num int) string {
	defer u.mu.RUnlock()
	u.mu.RLock()
	return u.questions[id][num]
}

func (u *UserQuestions) GetUserAns(id uint64, num int) string {
	defer u.mu.RUnlock()
	u.mu.RLock()
	return u.answers[id][num]
}

func (u *UserQuestions) DeleteUser(id uint64) error {
	defer u.mu.Unlock()
	u.mu.Lock()
	_, ok := u.questions[id]
	if !ok {
		return status.Errorf(codes.NotFound, "not found user")
	}
	delete(u.questions, id)
	delete(u.answers, id)
	delete(u.numQuestions, id)
	delete(u.correctAnsCount, id)
	return nil
}

func (u *UserQuestions) GetCountAskedQuestions(id uint64) int {
	defer u.mu.RUnlock()
	u.mu.RLock()
	return u.numQuestions[id]
}

func (u *UserQuestions) IsEndUserQuestions(id uint64) bool {
	defer u.mu.RUnlock()
	u.mu.RLock()
	return u.numQuestions[id] == len(u.questions[id])
}

func (u *UserQuestions) IncUserAskedQuestions(id uint64) {
	defer u.mu.Unlock()
	u.mu.Lock()
	u.numQuestions[id]++
}

func (u *UserQuestions) IncUserCorrectAns(id uint64) {
	defer u.mu.Unlock()
	u.mu.Lock()
	u.correctAnsCount[id]++
}

func (u *UserQuestions) GetCorrectAnsUser(id uint64) int {
	defer u.mu.RUnlock()
	u.mu.RLock()
	return u.correctAnsCount[id]
}
