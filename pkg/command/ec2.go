package command

import (
	"errors"
	"fmt"

	"github.com/nlopes/slack"
	"github.com/odg0318/aws-slack-bot/pkg/aws"
	"github.com/odg0318/aws-slack-bot/pkg/context"
	"github.com/odg0318/aws-slack-bot/pkg/util"
)

var (
	errorEc2InvalidParams = errors.New("InvalidParams! `@ec2 [instance-id]`")
)

type Ec2Command struct {
	context    *context.Context
	channel    string
	instanceId string
}

func (c *Ec2Command) Run() error {
	client := c.context.GetClient()
	cfg := c.context.GetConfig()

	sessions, err := aws.NewSessions(cfg)
	if err != nil {
		return err
	}

	attachments := []slack.Attachment{}
	response, err := aws.FindEc2Ip(sessions, c.instanceId)
	if err != nil {
		return err
	}

	for _, instance := range response.Instances {
		attachment := slack.Attachment{
			Text:       fmt.Sprintf("`instance-id`: %v | `private-id`: %v | `public-ip`: %v", instance.ID, instance.PrivateIp, instance.PublicIp),
			Color:      "#12d34d",
			MarkdownIn: []string{"text"},
		}
		attachments = append(attachments, attachment)
	}

	util.SendAttatchment(client, c.channel, attachments)

	return nil
}

func newEc2Command(ctx *context.Context, channel string, params []string) (*Ec2Command, error) {
	if len(params) != 1 {
		return nil, errorEc2InvalidParams
	}

	return &Ec2Command{ctx, channel, params[0]}, nil
}
