package cache

/*
import (
	"context"
	"sync"
	"time"
)

type UserTimeout struct {
	mu           sync.RWMutex
	usersTimeout map[uint64]context.Context
}

func (u *UserTimeout) SetTime(ctx context.Context, id int) {
	defer u.mu.Unlock()
	u.mu.Lock()
	u.usersTimeout[id] = time
}

func (u *UserTimeout) IsUserTimeoutFinished(id uint64) bool {
	defer u.mu.RUnlock()
	u.mu.Lock()
	return true
}

*/
