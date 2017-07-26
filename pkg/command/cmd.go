package command

import (
	"strings"

	"github.com/odg0318/aws-slack-bot/pkg/context"
)

var (
	cmdEc2   = "@ec2"
	cmdDummy = "@dummy"
)

type Command interface {
	Run() error
}

func NewCommand(ctx *context.Context, text, channel string) Command {
	tokens := strings.Split(text, " ")
	cmd := tokens[0]
	params := tokens[1:]

	switch cmd {
	case cmdEc2:
		return newEc2Command(ctx, channel, params)
	case cmdDummy:
		return newDummyCommand(ctx, channel, params)
	default:
		return nil
	}
}
