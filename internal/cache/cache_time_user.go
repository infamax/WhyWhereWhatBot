package cache

import (
	"sync"
	"time"
)

type TimeUsers struct {
	mu        sync.RWMutex
	usersTime map[uint64]*time.Timer
}

func NewUserTimer() *TimeUsers {
	return &TimeUsers{
		usersTime: make(map[uint64]*time.Timer, 0),
	}
}

func (t *TimeUsers) SetTimerUser(id uint64, duration time.Duration) {
	defer t.mu.Unlock()
	t.mu.Lock()
	t.usersTime[id] = time.NewTimer(duration)
}

func (t *TimeUsers) IsUserTimeCompleted(id uint64) bool {
	defer t.mu.Unlock()
	t.mu.Lock()
	select {
	case <-t.usersTime[id].C:
		return true
	default:
		return false
	}
}
