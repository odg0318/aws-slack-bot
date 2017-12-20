package util

import (
	"github.com/nlopes/slack"
)

const (
	COLOR_SUCCESS = "#4dd14d"
	COLOR_FAIL    = "#FF0000"
)

var (
	COLOR_SUCCESS_LIST = []string{"#22f98d", "#00b95d", "#368992"}
)

func SendAttatchment(client *slack.Client, channel string, text string, attachments []slack.Attachment) {
	params := slack.PostMessageParameters{
		Attachments: attachments,
	}

	client.PostMessage(channel, text, params)
}

func SendError(client *slack.Client, channel string, err error) {
	params := slack.PostMessageParameters{
		Username: "awsbot",
		Attachments: []slack.Attachment{
			{
				Text:       err.Error(),
				Color:      COLOR_FAIL,
				MarkdownIn: []string{"text"},
			},
		},
	}

	client.PostMessage(channel, "", params)
}
