package repository

import (
	"github.com/AZRV17/Course_client/internal/domain"
	"github.com/AZRV17/Course_client/pkg/storage/inmem"
)

type CoinRepositoryInterface interface {
	GetAll() []domain.Coin
	Get(string) (domain.Coin, bool)
	Set(domain.Coin)
}

type Repository struct {
	Storage        *inmem.InMemoryCache
	CoinRepository CoinRepositoryInterface
}

func NewRepository(storage *inmem.InMemoryCache) *Repository {
	return &Repository{
		Storage:        storage,
		CoinRepository: NewCoinRepository(storage),
	}
}
