package receiveMessage

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"xlab-feishu-robot/internal/pkg"

	larkcore "github.com/larksuite/oapi-sdk-go/v3/core"
	larkim "github.com/larksuite/oapi-sdk-go/v3/service/im/v1"
	"github.com/sirupsen/logrus"
)

func group(event *larkim.P2MessageReceiveV1) {
	messageType := *event.Event.Message.MessageType
	switch strings.ToUpper(messageType) {
	case "TEXT":
		groupTextMessage(event)
	default:
		logrus.WithFields(logrus.Fields{"message type": messageType}).Warn("Receive group message, but this type is not supported")
	}
}

func groupTextMessage(event *larkim.P2MessageReceiveV1) {
	// get chatid
	chatid := *event.Event.Message.ChatId
	req := larkim.NewGetChatMembersReqBuilder().
		ChatId(chatid).
		Build()

	// 发起请求
	resp, err := pkg.Cli.Im.ChatMembers.Get(context.Background(), req, larkcore.WithUserAccessToken("u-cUOEALQFF4uGgwPfwLBt3lk12CRN01VHr200g0s00Kn."))

	// 处理错误
	if err != nil {
		fmt.Println(err)
		return
	}

	// 服务端错误处理
	if !resp.Success() {
		fmt.Println(resp.Code, resp.Msg, resp.RequestId())
		return
	}

	// 业务处理
	// fmt.Println(larkcore.Prettify(resp))
	// fmt.Println(resp.Data.Items)

	var items []struct {
		// MemberIdType string `json:"MemberIdType"`
		// MemberId     string `json:"MemberId"`
		// Name         string `json:"Name"`
		// TenantKey    string `json:"TenantKey"`
		MemberIdType *string `json:"member_id_type,omitempty"` // 成员的用户 ID 类型，与查询参数中的 member_id_type 相同。取值为：`open_id`、`user_id`、`union_id`其中之一。
		MemberId     *string `json:"member_id,omitempty"`      // 成员的用户ID，ID值与查询参数中的 member_id_type 对应。;;不同 ID 的说明参见 [用户相关的 ID 概念](https://open.feishu.cn/document/home/user-identity-introduction/introduction)
		Name         *string `json:"name,omitempty"`           // 名字
		TenantKey    *string `json:"tenant_key,omitempty"`     // 租户Key，为租户在飞书上的唯一标识，用来换取对应的tenant_access_token，也可以用作租户在应用中的唯一标识
	}
	itemsJSON, err := json.Marshal(resp.Data.Items)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}
	err = json.Unmarshal(itemsJSON, &items)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}
	var head string
	var b GBoard
	for _, item := range items {
		// fmt.Println("Name:", item.Name)
		// fmt.Print(*event.Event.Sender.SenderId.OpenId)
		// fmt.Print(item.MemberIdType)
		// fmt.Print(item.TenantKey)
		// fmt.Print(item.MemberId)
		if *event.Event.Sender.SenderId.OpenId == *item.MemberId {
			head = "<at user_id=\"" + *item.MemberId + "\">" + *item.Name + "</at>\n"
			_, exist := MG[*item.MemberId]
			if !exist {
				MG[*item.MemberId] = GNew()
				b = MG[*item.MemberId]
			} else {
				b = MG[*item.MemberId]
			}
			// fmt.Println("head:", head)
			break
		}
	}

	// get the pure text message
	content := *event.Event.Message.Content
	content = strings.TrimSuffix(strings.TrimPrefix(content, "{\"text\":\""), "\"}")
	// 在群组中，消息内容的前面往往会有一个@机器人的字符串，需要去掉
	content = content[strings.Index(content, " ")+1:]
	event.Event.Message.Content = &content
	logrus.WithFields(logrus.Fields{"message content": content}).Info("Receive group TEXT message")

	switch content {
	case "help":
		groupHelpMenu(event)
	case "initial":
		groupNewGame(event, head, b)
	case "上":
		b.Move(event, content)
		b.AddElement()
		b.GDisplay(event, head)
		if b.IsSuccess() {
			text := head + "You win!"
			SendMessage(UserOpenId, *event.Event.Sender.SenderId.OpenId, Text, text)
		}
		if b.IsOver() {
			text := head + "You lose!\n"
			// SendMessage(UserOpenId, *event.Event.Sender.SenderId.OpenId, Text, text)
			sum, score := g.CountScore()
			text += "Your score is " + fmt.Sprintf("%d", score) + " and the sum is " + fmt.Sprintf("%d", sum)
			SendMessage(GroupChatId, *event.Event.Message.ChatId, Text, text)
		}
	case "下":
		b.Move(event, content)
		b.AddElement()
		b.GDisplay(event, head)
		if b.IsSuccess() {
			text := head + "You win!"
			SendMessage(UserOpenId, *event.Event.Sender.SenderId.OpenId, Text, text)
		}
		if b.IsOver() {
			text := head + "You lose!\n"
			// SendMessage(UserOpenId, *event.Event.Sender.SenderId.OpenId, Text, text)
			sum, score := g.CountScore()
			text += "Your score is " + fmt.Sprintf("%d", score) + " and the sum is " + fmt.Sprintf("%d", sum)
			SendMessage(GroupChatId, *event.Event.Message.ChatId, Text, text)
		}
	case "左":
		b.Move(event, content)
		b.AddElement()
		b.GDisplay(event, head)
		if b.IsSuccess() {
			text := head + "You win!"
			SendMessage(UserOpenId, *event.Event.Sender.SenderId.OpenId, Text, text)
		}
		if b.IsOver() {
			text := head + "You lose!\n"
			// SendMessage(UserOpenId, *event.Event.Sender.SenderId.OpenId, Text, text)
			sum, score := g.CountScore()
			text += "Your score is " + fmt.Sprintf("%d", score) + " and the sum is " + fmt.Sprintf("%d", sum)
			SendMessage(GroupChatId, *event.Event.Message.ChatId, Text, text)
		}
	case "右":
		b.Move(event, content)
		b.AddElement()
		b.GDisplay(event, head)
		if b.IsSuccess() {
			text := head + "You win!"
			SendMessage(UserOpenId, *event.Event.Sender.SenderId.OpenId, Text, text)
		}
		if b.IsOver() {
			text := head + "You lose!\n"
			// SendMessage(UserOpenId, *event.Event.Sender.SenderId.OpenId, Text, text)
			sum, score := g.CountScore()
			text += "Your score is " + fmt.Sprintf("%d", score) + " and the sum is " + fmt.Sprintf("%d", sum)
			SendMessage(GroupChatId, *event.Event.Message.ChatId, Text, text)
		}
	default:
		logrus.WithFields(logrus.Fields{"message content": content}).Warn("Receive group TEXT message, but this content does not have a handler")
	}
}
