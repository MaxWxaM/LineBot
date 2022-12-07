package controller

import (
	"net/http"

	"github.com/MaxWxaM/linebot/internal/app/model/input"

	"github.com/MaxWxaM/linebot/internal/app/service"
	"github.com/gin-gonic/gin"
)

type MsgController struct {
	msgService service.IMsgService
}

func NewMsgController(msgService service.IMsgService) *MsgController {
	return &MsgController{
		msgService: msgService,
	}
}

func (ctl *MsgController) PushMsg(ctx *gin.Context) {
	msg := input.Msg{}
	if err := ctx.ShouldBindJSON(&msg); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctl.msgService.PushMsg(ctx, msg)
	ctx.String(http.StatusOK, "Connect success: "+msg.Header)
}
