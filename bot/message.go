package bot

import (
	"fmt"
	"strings"

	"github.com/nlopes/slack"
	"github.com/odg0318/aws-slack-bot/command"
	"github.com/odg0318/aws-slack-bot/context"
	"github.com/odg0318/aws-slack-bot/util"
)

type Message struct {
	context *context.Context
	channel string
	user    string
	text    string
}

func (m *Message) Skip() bool {
	if len(m.user) == 0 {
		return true
	}

	tokens := strings.Split(m.text, " ")
	if len(tokens) == 0 {
		return true
	}

	signal := fmt.Sprintf("<@%s>", m.context.GetBotInfo().ID)

	if tokens[0] != signal {
		return true
	}
	return false
}

func (m *Message) Run() error {
	cmd, err := command.NewCommand(m.context, m.text, m.channel)
	if err != nil {
		client := m.context.GetClient()
		util.SendError(client, m.channel, err)
		return err
	}

	if cmd == nil {
		return nil
	}

	err = cmd.Run()
	if err != nil {
		client := m.context.GetClient()
		util.SendError(client, m.channel, err)
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
