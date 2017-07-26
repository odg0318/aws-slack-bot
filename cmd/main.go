package main

import (
	"os"

	"github.com/odg0318/aws-slack-bot/pkg/bot"
	"github.com/odg0318/aws-slack-bot/pkg/config"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "Aws Slack Bot"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "config",
			EnvVar: "AWS_SLACK_BOT_CONFIG",
			Usage:  "Configuration file path",
		},
	}
	app.Action = func(ctx *cli.Context) error {
		cfg, err := config.NewConfigFromCli(ctx)
		if err != nil {
			return err
		}

		err = cfg.Validate()
		if err != nil {
			return err
		}

		bot := bot.NewBot(cfg)
		bot.Run()
		return nil
	}

	app.Run(os.Args)
}
