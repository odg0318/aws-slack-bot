package command

import (
	"errors"
	"flag"
	"fmt"

	"github.com/nlopes/slack"
	"github.com/odg0318/aws-slack-bot/aws"
	"github.com/odg0318/aws-slack-bot/context"
	"github.com/odg0318/aws-slack-bot/util"
)

var (
	errorOpsworksHelp = errors.New("*Help* `@botname opsworks [-name=instance-name]`")
)

type OpsworksCommand struct {
	context      *context.Context
	channel      string
	instanceName string
	region       string
}

func (c *OpsworksCommand) Parse(params []string) error {
	fset := flag.NewFlagSet("", flag.ContinueOnError)
	fset.StringVar(&c.instanceName, "name", "", "instance-name")
	fset.StringVar(&c.region, "region", "", "region")

	err := fset.Parse(params)
	if err != nil {
		return errorOpsworksHelp
	}

	if len(c.instanceName) == 0 {
		return errorOpsworksHelp
	}

	return nil
}

func (c *OpsworksCommand) Run() error {
	client := c.context.GetClient()
	cfg := c.context.GetConfig()

	sessions, err := aws.NewSessions(cfg)
	if err != nil {
		return err
	}

	attachments := []slack.Attachment{}
	response, err := aws.FindOpsworksIp(sessions, c.instanceName)
	if err != nil {
		return err
	}

	for _, instance := range response.Instances {
		fields := []slack.AttachmentField{
			{
				Title: "name",
				Value: instance.Name,
				Short: true,
			},
			{
				Title: "id",
				Value: instance.ID,
				Short: true,
			},
			{
				Title: "private ip",
				Value: instance.PrivateIp,
				Short: true,
			},
			{
				Title: "public ip",
				Value: instance.PublicIp,
				Short: true,
			},
		}
		attachment := slack.Attachment{
			Pretext:    fmt.Sprintf("*`%s` is found.*", instance.ID),
			Fields:     fields,
			Color:      util.COLOR_SUCCESS,
			MarkdownIn: []string{"pretext", "text"},
		}
		attachments = append(attachments, attachment)
	}

	util.SendAttatchment(client, c.channel, "", attachments)

	return nil
}

func newOpsworksCommand(ctx *context.Context, channel string, params []string) (*OpsworksCommand, error) {
	c := &OpsworksCommand{
		context: ctx,
		channel: channel,
	}

	err := c.Parse(params)
	if err != nil {
		return nil, err
	}

	return c, nil
}
