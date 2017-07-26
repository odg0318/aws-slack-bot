package bot

import (
	"github.com/nlopes/slack"
	"github.com/odg0318/aws-slack-bot/pkg/command"
	"github.com/odg0318/aws-slack-bot/pkg/context"
)

type Message struct {
	context *context.Context
	channel string
	user    string
	text    string
}

func (m *Message) Mine() bool {
	return len(m.user) == 0
}

func (m *Message) Run() error {
	cmd := command.NewCommand(m.context, m.text, m.channel)
	if cmd == nil {
		return nil
	}

	err := cmd.Run()
	if err != nil {
		return err
	}

	return nil
}

func NewMessage(ctx *context.Context, ev *slack.MessageEvent) *Message {
	return &Message{
		context: ctx,
		channel: ev.Channel,
		user:    ev.User,
		text:    ev.Text,
	}
}
