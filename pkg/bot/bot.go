package bot

import (
	"errors"
	"log"
	"os"

	"github.com/nlopes/slack"
	"github.com/odg0318/aws-slack-bot/pkg/config"
	"github.com/odg0318/aws-slack-bot/pkg/context"
)

var (
	errorInvalidCredential = errors.New("InvalidCredential")
)

type Bot struct {
	config *config.Config
	client *slack.Client
	logger *log.Logger
}

func (b *Bot) Run() error {
	rtm := b.client.NewRTM()
	go rtm.ManageConnection()

	for msg := range rtm.IncomingEvents {
		switch ev := msg.Data.(type) {
		case *slack.ConnectedEvent:
			b.logger.Print("Connected.")
		case *slack.MessageEvent:
			b.onMessageEvent(ev)
		case *slack.InvalidAuthEvent:
			return errorInvalidCredential
		default:
			//ignore
		}
	}

	return nil
}

func (b *Bot) onMessageEvent(ev *slack.MessageEvent) {
	ctx := context.NewContext()
	ctx.SetClient(b.client)
	m := NewMessage(ctx, ev)

	b.logger.Printf("[%s] %s\n", m.channel, m.text)

	if m.Mine() {
		return
	}

	err := m.Run()
	if err != nil {
		b.logger.Print(err)
	}
}

func NewBot(cfg *config.Config) *Bot {
	logger := log.New(os.Stdout, "slack-bot: ", log.Lshortfile|log.LstdFlags)
	slack.SetLogger(logger)

	client := slack.New(cfg.Slack.Token)
	client.SetDebug(cfg.Debug)

	logger = log.New(os.Stdout, "aws-bot: ", log.Lshortfile|log.LstdFlags)
	bot := &Bot{
		config: cfg,
		client: client,
		logger: logger,
	}

	return bot
}
