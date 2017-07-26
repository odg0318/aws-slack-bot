package command

import (
	"github.com/nlopes/slack"
	"github.com/odg0318/aws-slack-bot/pkg/context"
)

type Ec2Command struct {
	context *context.Context
	channel string
	params  []string
}

func (c *Ec2Command) Run() error {
	client := c.context.GetClient()

	params := slack.PostMessageParameters{}
	attachment := slack.Attachment{
		Pretext: "some pretext",
		Text:    "some text",
	}
	params.Attachments = []slack.Attachment{attachment}

	client.PostMessage(c.channel, "", params)

	return nil
}

func newEc2Command(ctx *context.Context, channel string, params []string) *Ec2Command {
	return &Ec2Command{ctx, channel, params}
}
