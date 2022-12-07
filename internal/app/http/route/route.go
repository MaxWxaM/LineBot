package route

import (
	"github.com/MaxWxaM/linebot/internal/app/http/controller"

	"github.com/gin-gonic/gin"
)

func NewRoute(msgController *controller.MsgController) *gin.Engine {
	r := gin.Default()
	r.POST("/msg", msgController.PushMsg)
	return r
}
