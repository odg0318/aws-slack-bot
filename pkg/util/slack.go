package util

import (
	"github.com/nlopes/slack"
)

func SendAttatchment(client *slack.Client, channel string, attachments []slack.Attachment) {
	params := slack.PostMessageParameters{
		Attachments: attachments,
	}

	client.PostMessage(channel, "", params)
}

func SendError(client *slack.Client, channel string, err error) {
	params := slack.PostMessageParameters{
		Username: "awsbot",
		Attachments: []slack.Attachment{
			{
				Text:       err.Error(),
				Color:      "#ff0000",
				MarkdownIn: []string{"text"},
			},
		},
	}

	client.PostMessage(channel, "", params)
}
