package repository

import (
	"github.com/AZRV17/Course_client/internal/domain"
	"github.com/AZRV17/Course_client/pkg/storage/inmem"
)

type CoinRepository struct {
	storage *inmem.InMemoryCache
}

func NewCoinRepository(storage *inmem.InMemoryCache) *CoinRepository {
	return &CoinRepository{
		storage: storage,
	}
}

func (c *CoinRepository) Set(coin domain.Coin) {
	c.storage.Mu.Lock()
	defer c.storage.Mu.Unlock()
	c.storage.Coins[coin.ID] = coin
}

func (c *CoinRepository) Get(coinID string) (domain.Coin, bool) {
	c.storage.Mu.RLock()
	defer c.storage.Mu.RUnlock()
	coin, exists := c.storage.Coins[coinID]
	return coin, exists
}

func (c *CoinRepository) GetAll() []domain.Coin {
	c.storage.Mu.RLock()
	defer c.storage.Mu.RUnlock()

	coins := make([]domain.Coin, 0, len(c.storage.Coins))
	for _, v := range c.storage.Coins {
		coins = append(coins, v)
	}
	return coins
}
