package receiveMessage

import (
	"context"

	larkcard "github.com/larksuite/oapi-sdk-go/v3/card"
	larkim "github.com/larksuite/oapi-sdk-go/v3/service/im/v1"
	"github.com/sirupsen/logrus"
)

// Receive dispatch message according to chat type
func Receive(_ context.Context, event *larkim.P2MessageReceiveV1) error {
	chatType := *event.Event.Message.ChatType
	switch chatType {
	case "p2p":
		p2p(event)
	case "group":
		group(event)
	// add more chat type here if needed

	default:
		logrus.WithFields(logrus.Fields{"chat type": chatType}).Warn("Receive message, but this chat type is not supported")
	}

	return nil
}

// Receive user message sending to the robot, a demo
func ReceiveCard(_ context.Context, event *larkcard.CardAction) (interface{}, error) {
	// 创建 http body
	body := make(map[string]interface{})
	body["content"] = "hello"

	i18n := make(map[string]string)
	i18n["zh_cn"] = "你好"
	i18n["en_us"] = "hello"
	i18n["ja_jp"] = "こんにちは"
	body["i18n"] = i18n

	// 创建自定义消息：http状态码，body内容
	resp := &larkcard.CustomResp{
		StatusCode: 400,
		Body:       body,
	}

	return resp, nil
}
