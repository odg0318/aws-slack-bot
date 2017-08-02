package bot

import (
	"errors"
	"log"
	"os"

	"github.com/nlopes/slack"
	"github.com/odg0318/aws-slack-bot/config"
	"github.com/odg0318/aws-slack-bot/context"
	"github.com/odg0318/aws-slack-bot/util"
)

var (
	errorInvalidCredential = errors.New("InvalidCredential")
)

type Bot struct {
	info        *slack.UserDetails
	config      *config.Config
	client      *slack.Client
	logger      *log.Logger
	lastChannel string
}

func (b *Bot) Run() error {
	defer func() {
		if r := recover(); r != nil {
			util.SendError(b.client, b.lastChannel, r.(error))

			b.logger.Printf("Panic! %+v\n", r)
			b.Run()
		}
	}()

	rtm := b.client.NewRTM()
	go rtm.ManageConnection()

	for msg := range rtm.IncomingEvents {
		switch ev := msg.Data.(type) {
		case *slack.ConnectedEvent:
			b.onConnectedEvent(ev)
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

func (b *Bot) onConnectedEvent(ev *slack.ConnectedEvent) {
	b.info = ev.Info.User
	b.logger.Printf("Connected. %s", b.info.Name)
}

func (b *Bot) onMessageEvent(ev *slack.MessageEvent) {
	ctx := context.NewContext()
	ctx.SetClient(b.client)
	ctx.SetConfig(b.config)
	ctx.SetBotInfo(b.info)

	m := NewMessage(ctx, ev)

	b.logger.Printf("[%s] %s\n", m.channel, m.text)

	if m.Skip() {
		return
	}

	b.lastChannel = m.channel

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
