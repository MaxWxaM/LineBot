package service

import (
	"net/http"

	"github.com/MaxWxaM/linebot/internal/app/model/client"
	"github.com/MaxWxaM/linebot/internal/app/model/input"
	"github.com/MaxWxaM/linebot/internal/app/repo"
	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/linebot"
)

type IMsgService interface {
	PushMsg(ctx *gin.Context, msg input.Msg)
}

type MsgService struct {
	msgRepo repo.IMsgRepository
	bot     *client.CustomizeLineBotClient
}

func NewMsgService(msgRepo repo.IMsgRepository, bot *client.CustomizeLineBotClient) IMsgService {
	return &MsgService{
		msgRepo: msgRepo,
		bot:     bot,
	}
}

func (service MsgService) PushMsg(ctx *gin.Context, msg input.Msg) {
	userId := service.bot.UserId
	_, err := service.bot.Client.PushMessage(userId, linebot.NewTextMessage("Hello, "+msg.Header+" "+msg.Content)).Do()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	err = service.msgRepo.LogMsg(ctx, msg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
}
