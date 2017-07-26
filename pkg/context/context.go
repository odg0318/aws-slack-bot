package context

import (
	"errors"

	"github.com/nlopes/slack"
	"github.com/odg0318/aws-slack-bot/pkg/config"
)

var (
	errorInvalidKeyContext = errors.New("InvalidKeyContext")
	errorExistKeyContext   = errors.New("ExistKeyContext")

	keyClient  = "client"
	keyConfig  = "config"
	keyBotInfo = "botinfo"
)

type Context struct {
	data map[string]interface{}
}

func (c *Context) Set(key string, val interface{}) error {
	if _, ok := c.data[key]; ok {
		return errorExistKeyContext
	}

	c.data[key] = val

	return nil
}

func (c *Context) Get(key string) (interface{}, error) {
	val, ok := c.data[key]
	if ok == false {
		return nil, errorInvalidKeyContext
	}
	return val, nil
}

func (c *Context) SetClient(client *slack.Client) error {
	return c.Set(keyClient, client)
}

func (c *Context) GetClient() *slack.Client {
	val, _ := c.Get(keyClient)

	return val.(*slack.Client)
}

func (c *Context) SetConfig(cfg *config.Config) error {
	return c.Set(keyConfig, cfg)
}

func (c *Context) GetConfig() *config.Config {
	val, _ := c.Get(keyConfig)

	return val.(*config.Config)
}

func (c *Context) SetBotInfo(cfg *slack.UserDetails) error {
	return c.Set(keyBotInfo, cfg)
}

func (c *Context) GetBotInfo() *slack.UserDetails {
	val, _ := c.Get(keyBotInfo)

	return val.(*slack.UserDetails)
}

func NewContext() *Context {
	return &Context{
		data: map[string]interface{}{},
	}
}
