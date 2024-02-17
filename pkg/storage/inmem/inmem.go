package inmem

import (
	"github.com/AZRV17/Course_client/internal/domain"
	"sync"
)

type InMemoryCache struct {
	Mu    sync.RWMutex
	Coins map[string]domain.Coin
}

func NewInMemoryCache() *InMemoryCache {
	return &InMemoryCache{
		Coins: make(map[string]domain.Coin),
	}
}
