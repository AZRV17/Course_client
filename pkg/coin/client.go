package coin

import (
	"encoding/json"
	"github.com/AZRV17/Course_client/internal/domain"
	"github.com/AZRV17/Course_client/internal/repository"
	"io"
	"log"
	"net/http"
	"time"
)

type Client struct {
	httpClient *http.Client
	url        string
}

func NewClient(url string, timeout time.Duration) *Client {
	return &Client{
		httpClient: &http.Client{
			Timeout: timeout,
		},
		url: url,
	}
}

func (c *Client) FetchData(repo *repository.Repository) error {
	response, err := c.httpClient.Get(c.url)
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(response.Body)

	var coins []domain.Coin

	data, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}

	if err = json.Unmarshal(data, &coins); err != nil {
		return err
	}

	for _, coin := range coins {
		repo.CoinRepository.Set(coin)
	}

	return nil
}
