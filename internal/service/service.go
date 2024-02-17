package service

import (
	"github.com/AZRV17/Course_client/internal/domain"
	"github.com/AZRV17/Course_client/internal/repository"
)

type CoinServiceInterface interface {
	GetAll() []domain.Coin
	Set(domain.Coin)
	Get(string) (domain.Coin, bool)
}

type Service struct {
	Repository  *repository.Repository
	CoinService CoinServiceInterface
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Repository:  repo,
		CoinService: NewCoinService(repo),
	}
}
