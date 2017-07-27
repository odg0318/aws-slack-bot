package aws

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/odg0318/aws-slack-bot/config"
)

type Sessions []*session.Session

func NewSessions(cfg *config.Config) (Sessions, error) {
	sessions := Sessions{}
	for _, a := range cfg.Aws {
		session, err := newSession(a.Region, a.AccessKey, a.SecretAccessKey)
		if err != nil {
			return nil, err
		}
		sessions = append(sessions, session)
	}

	return sessions, nil
}

func newSession(region, accessKey, secretAccessKey string) (*session.Session, error) {
	awsConfig := &aws.Config{
		Region:      aws.String(region),
		Credentials: credentials.NewStaticCredentials(accessKey, secretAccessKey, ""),
	}

	options := session.Options{
		Config: *awsConfig,
	}

	sess, err := session.NewSessionWithOptions(options)
	if err != nil {
		return nil, err
	}

	return sess, nil
}
