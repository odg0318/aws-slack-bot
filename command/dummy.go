package command

import (
	"strings"

	"github.com/nlopes/slack"

	"github.com/odg0318/aws-slack-bot/context"
	"github.com/odg0318/aws-slack-bot/util"
)

type DummyCommand struct {
	context *context.Context
	channel string
	cmds    []string
}

func (c *DummyCommand) Parse(params []string) error {
	c.cmds = allCmds
	return nil
}

func (c *DummyCommand) Run() error {
	client := c.context.GetClient()

	attachment := slack.Attachment{
		Text:  strings.Join(c.cmds, " "),
		Color: "#ff0000",
	}
	attachments := []slack.Attachment{attachment}

	util.SendAttatchment(client, c.channel, "Hi! I'm aws slack bot. There are following commands.", attachments)

	return nil
}

func newDummyCommand(ctx *context.Context, channel string, params []string) (*DummyCommand, error) {
	c := &DummyCommand{
		context: ctx,
		channel: channel,
	}

	err := c.Parse(params)
	if err != nil {
		return nil, err
	}

	return c, nil
}
