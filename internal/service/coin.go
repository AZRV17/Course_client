package service

import (
	"github.com/AZRV17/Course_client/internal/domain"
	"github.com/AZRV17/Course_client/internal/repository"
)

type CoinService struct {
	Repository *repository.Repository
}

func NewCoinService(repo *repository.Repository) *CoinService {
	return &CoinService{
		Repository: repo,
	}
}

func (c CoinService) GetAll() []domain.Coin {
	return c.Repository.CoinRepository.GetAll()
}

func (c CoinService) Set(coin domain.Coin) {
	c.Repository.CoinRepository.Set(coin)
}

func (c CoinService) Get(coinID string) (domain.Coin, bool) {
	return c.Repository.CoinRepository.Get(coinID)
}
