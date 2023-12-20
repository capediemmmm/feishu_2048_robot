package receiveMessage

import (
	"context"
	"encoding/json"
	"xlab-feishu-robot/internal/pkg"

	larkim "github.com/larksuite/oapi-sdk-go/v3/service/im/v1"
	"github.com/sirupsen/logrus"
)

type MsgReceiverType string

const (
	UserOpenId  MsgReceiverType = "open_id"
	UserUnionId MsgReceiverType = "union_id"
	UserUserId  MsgReceiverType = "user_id"
	UserEmail   MsgReceiverType = "email"
	GroupChatId MsgReceiverType = "chat_id"
)

type MsgContentType string

const (
	Text        MsgContentType = "text"
	Interactive MsgContentType = "interactive"
)

const (
	_rows = 4
	_cols = 4
)

type Board interface {
	Display(event *larkim.P2MessageReceiveV1)
	AddElement()
	Move(event *larkim.P2MessageReceiveV1, content string)
	IsOver() bool
	CountScore() (int, int)
	IsSuccess() bool
}

type board struct {
	matrix  [][]int
	success bool
	newRow  int
	newCol  int
}

type GBoard interface {
	GDisplay(event *larkim.P2MessageReceiveV1, head string)
	AddElement()
	Move(event *larkim.P2MessageReceiveV1, content string)
	IsOver() bool
	CountScore() (int, int)
	IsSuccess() bool
}

var MG = map[string]GBoard{}

var g Board

// var G GBoard

// SendMessage Send a message to a person / chat group
func SendMessage(receiveIdType MsgReceiverType, receiveId string, msgType MsgContentType, msg string) {
	content := ""

	switch msgType {
	case Text:
		contentMap := make(map[string]string)
		contentMap["text"] = msg
		byteContent, err := json.Marshal(contentMap)
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"receive_id_type": string(receiveIdType),
				"receive_id":      receiveId,
				"msg_type":        string(msgType),
				"msg":             msg,
			}).Error("marshal text to json fail")
			return
		}
		content = string(byteContent)
	case Interactive:
		content = msg
	// add more message type here if needed
	default:
		logrus.WithField("msgType", msgType).Error("message type unsupported")
		return
	}

	req := larkim.NewCreateMessageReqBuilder().
		ReceiveIdType(string(receiveIdType)).
		Body(larkim.NewCreateMessageReqBodyBuilder().
			ReceiveId(receiveId).
			MsgType(string(msgType)).
			Content(content).
			Build()).
		Build()

	resp, err := pkg.Cli.Im.Message.Create(context.Background(), req)
	if err != nil {
		logrus.Error(err)
		return
	}

	if !resp.Success() {
		logrus.Error(resp.Code, resp.Msg)
		return
	}
}
