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
	errorEc2Help = errors.New("*Help* `@botname ec2 [-id=instance-id|-name=instance-name]`")
)

type Ec2Command struct {
	context      *context.Context
	channel      string
	instanceId   string
	instanceName string
	region       string
}

func (c *Ec2Command) Parse(params []string) error {
	fset := flag.NewFlagSet("", flag.ContinueOnError)
	fset.StringVar(&c.instanceId, "id", "", "instance-id")
	fset.StringVar(&c.instanceName, "name", "", "instance-name")
	fset.StringVar(&c.region, "region", "", "region")

	err := fset.Parse(params)
	if err != nil {
		return errorEc2Help
	}

	if len(c.instanceId) == 0 && len(c.instanceName) == 0 {
		return errorEc2Help
	}

	if len(c.instanceId) > 0 && len(c.instanceName) > 0 {
		return errorEc2Help
	}

	return nil
}

func (c *Ec2Command) Run() error {
	client := c.context.GetClient()
	cfg := c.context.GetConfig()

	sessions, err := aws.NewSessions(cfg)
	if err != nil {
		return err
	}

	attachments := []slack.Attachment{}
	response, err := aws.FindEc2Ip(sessions, c.instanceId, c.instanceName)
	if err != nil {
		return err
	}

	for i, instance := range response.Instances {
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
			Color:      util.COLOR_SUCCESS_LIST[i%3],
			MarkdownIn: []string{"pretext", "text"},
		}
		attachments = append(attachments, attachment)
	}

	util.SendAttatchment(client, c.channel, "", attachments)

	return nil
}

func newEc2Command(ctx *context.Context, channel string, params []string) (*Ec2Command, error) {
	c := &Ec2Command{
		context: ctx,
		channel: channel,
	}

	err := c.Parse(params)
	if err != nil {
		return nil, err
	}

	return c, nil
}
