package config

import (
	"github.com/spf13/viper"
	"time"
)

type Config struct {
	HTTP struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	}
	Client struct {
		Url     string        `yaml:"url"`
		Timeout time.Duration `yaml:"timeout"`
	}
}

func NewConfig(configPath string) (*Config, error) {
	cfg := Config{}

	viper.SetConfigFile(configPath)
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	cfg.HTTP.Host = viper.GetString("http.host")
	cfg.HTTP.Port = viper.GetString("http.port")
	cfg.Client.Url = viper.GetString("client.url")
	cfg.Client.Timeout = viper.GetDuration("client.timeout")

	return &cfg, nil
}
