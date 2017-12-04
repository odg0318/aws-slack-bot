FROM golang:1.8.3
WORKDIR /go/src/github.com/odg0318/aws-slack-bot
COPY . .

WORKDIR cmd/aws-slack-bot

RUN go-wrapper download
RUN go-wrapper install

CMD ["go-wrapper", "run"]
