package config

import (
	"gopkg.in/yaml.v2"
	"log"
)

// Временная структура для удобного парсинга yaml
type file struct {
	APIKeys struct {
		Telegram string `yaml:"telegram"`
		Site     string `yaml:"site"`
	} `yaml:"apiKeys"`
	Dsn        string `yaml:"dsn"`
	Moderators []struct {
		Name     string `yaml:"name"`
		Telegram string `yaml:"telegram"`
	} `yaml:"moderators"`
}

// Функция для парса конфига
func ParseConfig(fileBytes []byte) (*Config, error) {
	cf := file{}
	err := yaml.Unmarshal(fileBytes, &cf)
	if err != nil {
		return nil, err
	}

	c := Config{}

	log.Println("Parse api key = ", cf.APIKeys.Telegram)
	c.ApiKey.Telegram = cf.APIKeys.Telegram
	c.ApiKey.Site = cf.APIKeys.Site
	c.Moderators = make(map[UserId]string)
	c.Dsn = cf.Dsn

	for _, moderator := range cf.Moderators {
		c.Moderators[UserId(moderator.Telegram)] = moderator.Name
	}

	return &c, nil
}
