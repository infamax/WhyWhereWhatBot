package cache

import (
	"sync"
)

type UserGame struct {
	mu sync.RWMutex
	c  map[uint64]bool
}

func NewUserGame() *UserGame {
	return &UserGame{
		c: make(map[uint64]bool),
	}
}

func (c *UserGame) SetUserStartGame(id uint64) {
	defer c.mu.Unlock()
	c.mu.Lock()
	c.c[id] = true
}

func (c *UserGame) SetUserStopGame(id uint64) {
	defer c.mu.Unlock()
	c.mu.Lock()
	c.c[id] = false
}

func (c *UserGame) IsUserPlayGame(id uint64) bool {
	defer c.mu.RUnlock()
	c.mu.RLock()
	return c.c[id]
}

func (c *UserGame) GetUsersInGame() []uint64 {
	defer c.mu.RUnlock()
	c.mu.RLock()
	res := make([]uint64, 0)
	for key, value := range c.c {
		if value {
			res = append(res, key)
		}
	}
	return res
}
