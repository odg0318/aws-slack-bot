# aws-slack-bot
This slack bot helps you manage AWS resources via Slack.

Get Started
===========
1. Download package `go get -u github.com/odg0318/aws-slack-bot/cmd/aws-slack-bot`
2. Go to the package path `cd $GOPATH/odg0318/aws-slack-bot/cmd/aws-slack-bot`
3. Make sure that config.yml is valid. You can see an example in config.sample.yml
4. Build and run `go build && ./aws-slack-bot --config=/path/to/config.yml`
    
With Docker
===========
`Dockefile` is located in `docker` directory.

    $ make docker-build
    $ make docker-run
    
Availalbe Commands
==================
### EC2
#### Search by name
NOTE: Regular expression is available.

    @botname ec2 -name="ec2-name"
#### Search by id
    @botname ec2 -id="i-xxxx"
### Opsworks
NOTE: Name doesn't mean `Name` in tags but instance name in Opsworks. To search `Name` in tag, use `ec2` command.
#### Search by name
    @botname opsworks -name="instance-name"

Configuration
=============
```yaml
debug: false
bot:
  name: "BOT_NAME"
  emoji: "BOT_EMOJI"
slack:
  token: "SLACK_TOKEN_HERE"
aws:
  - region: "AWS_REGION"
    access_key: "AWS_ACCESS_KEY"
    secret_access_key: "AWS_SECRET_ACCESS_KEY"
```
Configuration is very simple but `aws` section needs some explaination.
To use multiple regions or accounts, `aws` section should be configured as an array including api key information.
```yaml
aws:
  - region: "us-east-1"
    access_key: "AWS_ACCESS_KEY"
    secret_access_key: "AWS_SECRET_ACCESS_KEY"
  - region: "ap-northeast-2"
    access_key: "AWS_ACCESS_KEY"
    secret_access_key: "AWS_SECRET_ACCESS_KEY"
```
