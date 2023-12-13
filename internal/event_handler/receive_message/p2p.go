package receiveMessage

import (
	"fmt"
	"strings"

	larkim "github.com/larksuite/oapi-sdk-go/v3/service/im/v1"
	"github.com/sirupsen/logrus"
)

func p2p(event *larkim.P2MessageReceiveV1) {
	messageType := *event.Event.Message.MessageType
	switch strings.ToUpper(messageType) {
	case "TEXT":
		p2pTextMessage(event)
	default:
		logrus.WithFields(logrus.Fields{"message type": messageType}).Warn("Receive p2p message, but this type is not supported")
	}
}

func p2pTextMessage(event *larkim.P2MessageReceiveV1) {
	// get the pure text message
	content := *event.Event.Message.Content
	content = strings.TrimSuffix(strings.TrimPrefix(content, "{\"text\":\""), "\"}")
	event.Event.Message.Content = &content
	logrus.WithFields(logrus.Fields{"message content": content}).Info("Receive p2p TEXT message")

	switch content {
	case "help":
		p2pHelpMenu(event)
	case "initial":
		p2pNewGame(event)
	case "上":
		g.Move(event, content)
		g.AddElement()
		g.Display(event)
		if g.IsSuccess() {
			SendMessage(UserOpenId, *event.Event.Sender.SenderId.OpenId, Text, "You win!")
		}
		if g.IsOver() {
			SendMessage(UserOpenId, *event.Event.Sender.SenderId.OpenId, Text, "You lose!")
			sum, score := g.CountScore()
			text := "Your score is " + fmt.Sprintf("%d", score) + " and the sum is " + fmt.Sprintf("%d", sum)
			SendMessage(UserOpenId, *event.Event.Sender.SenderId.OpenId, Text, text)
		}
	case "下":
		g.Move(event, content)
		g.AddElement()
		g.Display(event)
		if g.IsSuccess() {
			SendMessage(UserOpenId, *event.Event.Sender.SenderId.OpenId, Text, "You win!")
		}
		if g.IsOver() {
			SendMessage(UserOpenId, *event.Event.Sender.SenderId.OpenId, Text, "You lose!")
			sum, score := g.CountScore()
			text := "Your score is " + fmt.Sprintf("%d", score) + " and the sum is " + fmt.Sprintf("%d", sum)
			SendMessage(UserOpenId, *event.Event.Sender.SenderId.OpenId, Text, text)
		}
	case "左":
		g.Move(event, content)
		g.AddElement()
		g.Display(event)
		if g.IsSuccess() {
			SendMessage(UserOpenId, *event.Event.Sender.SenderId.OpenId, Text, "You win!")
		}
		if g.IsOver() {
			SendMessage(UserOpenId, *event.Event.Sender.SenderId.OpenId, Text, "You lose!")
			sum, score := g.CountScore()
			text := "Your score is " + fmt.Sprintf("%d", score) + " and the sum is " + fmt.Sprintf("%d", sum)
			SendMessage(UserOpenId, *event.Event.Sender.SenderId.OpenId, Text, text)
		}
	case "右":
		g.Move(event, content)
		g.AddElement()
		g.Display(event)
		if g.IsSuccess() {
			SendMessage(UserOpenId, *event.Event.Sender.SenderId.OpenId, Text, "You win!")
		}
		if g.IsOver() {
			SendMessage(UserOpenId, *event.Event.Sender.SenderId.OpenId, Text, "You lose!")
			sum, score := g.CountScore()
			text := "Your score is " + fmt.Sprintf("%d", score) + " and the sum is " + fmt.Sprintf("%d", sum)
			SendMessage(UserOpenId, *event.Event.Sender.SenderId.OpenId, Text, text)
		}
	default:
		logrus.WithFields(logrus.Fields{"message content": content}).Warn("Receive p2p TEXT message, but this content does not have a handler")
	}
}
