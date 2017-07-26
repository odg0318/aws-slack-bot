package command

import (
	"strings"

	"github.com/nlopes/slack"
	"github.com/odg0318/aws-slack-bot/pkg/context"
)

type DummyCommand struct {
	context *context.Context
	channel string
	params  []string
}

func (c *DummyCommand) Run() error {
	client := c.context.GetClient()

	params := slack.PostMessageParameters{}
	attachment := slack.Attachment{
		Pretext: "Dummy Command",
		Text:    strings.Join(c.params, " "),
	}
	params.Attachments = []slack.Attachment{attachment}

	client.PostMessage(c.channel, "", params)

	return nil
}

func newDummyCommand(ctx *context.Context, channel string, params []string) (*DummyCommand, error) {
	return &DummyCommand{ctx, channel, params}, nil
}
