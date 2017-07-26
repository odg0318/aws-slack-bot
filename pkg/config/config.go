package config

import (
	"errors"
	"io/ioutil"

	"github.com/urfave/cli"
	"gopkg.in/yaml.v2"
)

var (
	errorInvalidToken = errors.New("InvalidToken")
)

type Config struct {
	Debug bool `yaml:"debug"`
	Slack struct {
		Token string `yaml:"token"`
	} `yaml:"slack"`
	Aws []struct {
		AccessKey       string `yaml:"access_key"`
		SecretAccessKey string `yaml:"secret_access_key"`
	} `yaml:"aws"`
}

func (c *Config) Validate() error {
	if len(c.Slack.Token) == 0 {
		return errorInvalidToken
	}

	return nil
}

func NewConfigFromCli(ctx *cli.Context) (*Config, error) {
	cfgPath := ctx.String("config")
	data, err := ioutil.ReadFile(cfgPath)
	if err != nil {
		return nil, err
	}

	var cfg Config
	err = yaml.Unmarshal([]byte(data), &cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
