package command

import (
	"errors"
	"strings"

	"github.com/odg0318/aws-slack-bot/pkg/context"
)

var (
	cmdEc2   = "ec2"
	cmdDummy = "dummy"

	errorInvalidParams = errors.New("InvalidParams")
)

type Command interface {
	Run() error
}

func NewCommand(ctx *context.Context, text, channel string) (Command, error) {
	tokens := strings.Split(text, " ")
	cmd := tokens[1]
	params := tokens[2:]

	println(cmd, params)

	switch cmd {
	case cmdEc2:
		return newEc2Command(ctx, channel, params)
	case cmdDummy:
		return newDummyCommand(ctx, channel, params)
	default:
		return nil, nil
	}
}
