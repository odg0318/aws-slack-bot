package context

import (
	"errors"

	"github.com/nlopes/slack"
)

var (
	errorInvalidKeyContext = errors.New("InvalidKeyContext")
	errorExistKeyContext   = errors.New("ExistKeyContext")
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
	return c.Set("client", client)
}

func (c *Context) GetClient() *slack.Client {
	val, _ := c.Get("client")

	return val.(*slack.Client)
}

func NewContext() *Context {
	return &Context{
		data: map[string]interface{}{},
	}
}
