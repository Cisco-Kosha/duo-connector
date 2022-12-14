package config

import (
	"flag"
	"os"
)

type Config struct {
	personalAccessToken string
	duoSecurityURL      string
}

func Get() *Config {
	conf := &Config{}
	flag.StringVar(&conf.personalAccessToken, "AccessToken", os.Getenv("API_TOKEN_OR_KEY"), "Duo Access token or key")
	flag.StringVar(&conf.duoSecurityURL, "DuoSecurityURL", os.Getenv("DuoSecurityURL"), "Duo Security URL")
	flag.Parse()
	return conf
}

func (c *Config) GetPersonalAccessToken() string {
	return c.personalAccessToken
}

func (c *Config) GetDuoSecurityURL() string {
	return c.duoSecurityURL
}
