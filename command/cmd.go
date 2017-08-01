package command

import (
	"errors"
	"strings"

	"github.com/odg0318/aws-slack-bot/context"
)

var (
	allCmds     = []string{cmdEc2, cmdOpsworks}
	cmdEc2      = "ec2"
	cmdOpsworks = "opsworks"
	cmdDummy    = "dummy"

	errorInvalidParams = errors.New("InvalidParams")
)

type Command interface {
	Run() error
	Parse([]string) error
}

func NewCommand(ctx *context.Context, text, channel string) (Command, error) {
	tokens := strings.Split(text, " ")

	var cmd string
	var params []string

	if len(tokens) > 1 {
		cmd = tokens[1]
		params = tokens[2:]
	}

	switch cmd {
	case cmdEc2:
		return newEc2Command(ctx, channel, params)
	case cmdOpsworks:
		return newOpsworksCommand(ctx, channel, params)
	default:
		return newDummyCommand(ctx, channel, params)
	}
}
