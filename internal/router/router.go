package router

import (
	"xlab-feishu-robot/internal/config"
	"xlab-feishu-robot/internal/controller"
	receiveMessage "xlab-feishu-robot/internal/event_handler/receive_message"

	"github.com/gin-gonic/gin"
	sdkginext "github.com/larksuite/oapi-sdk-gin"
	larkcard "github.com/larksuite/oapi-sdk-go/v3/card"
	"github.com/larksuite/oapi-sdk-go/v3/event/dispatcher"
)

func Register(r *gin.Engine) {
	// example
	r.POST("/api/example", controller.Example)

	// register dispatcher
	handler := dispatcher.NewEventDispatcher(config.C.LarkConfig.VerificationToken, config.C.LarkConfig.EncryptKey).
		OnP2MessageReceiveV1(receiveMessage.Receive)
	// DO NOT CHANGE THIS LINE
	r.POST("/lark/event", sdkginext.NewEventHandlerFunc(handler))
	cardHandler := larkcard.NewCardActionHandler(config.C.LarkConfig.VerificationToken, config.C.LarkConfig.EncryptKey, receiveMessage.ReceiveCard)
	r.POST("/lark/2048", sdkginext.NewCardActionHandlerFunc(cardHandler))
}