package client

import "github.com/line/line-bot-sdk-go/linebot"

type CustomizeLineBotClient struct {
	*linebot.Client
	UserId string
}
