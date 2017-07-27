package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestErrorValidateBotName(t *testing.T) {
	c := Config{}

	err := c.Validate()

	assert.NotNil(t, err)
	assert.Equal(t, err, errorMustBotName)
}

func TestErrorValidateSlackToken(t *testing.T) {
	c := Config{
		Bot: struct {
			Name  string `yaml:"name"`
			Emoji string `yaml:"emoji"`
		}{"awsbot", ":emoji:"},
	}

	err := c.Validate()

	assert.NotNil(t, err)
	assert.Equal(t, err, errorMustSlackToken)
}
