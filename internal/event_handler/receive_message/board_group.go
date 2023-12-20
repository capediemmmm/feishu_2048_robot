package receiveMessage

import (
	"fmt"

	larkim "github.com/larksuite/oapi-sdk-go/v3/service/im/v1"
)

func groupHelpMenu(event *larkim.P2MessageReceiveV1) {
	SendMessage(GroupChatId, *event.Event.Message.ChatId, Text, "this is a group test string")
}

func groupNewGame(event *larkim.P2MessageReceiveV1, head string, b GBoard) {
	b = GNew()
	b.AddElement()
	b.AddElement()
	b.GDisplay(event, head)
}

func GNew() GBoard {
	matrix := make([][]int, 0)
	for i := 0; i < _rows; i++ {
		matrix = append(matrix, make([]int, _cols))
	}
	return &board{
		matrix: matrix,
	}
}

func (b *board) GDisplay(event *larkim.P2MessageReceiveV1, head string) {
	var table string
	table += head
	//b.matrix = getRandom()
	for i := 0; i < len(b.matrix); i++ {
		for i := 0; i < 40; i++ {
			table += "-"
		}
		table += "\n"
		table += "|"
		for j := 0; j < len(b.matrix[0]); j++ {
			table += fmt.Sprintf("%3s", " ")
			if b.matrix[i][j] == 0 {
				table += fmt.Sprintf("%-6s|", " ")
			} else {
				table += fmt.Sprintf("%-6d|", b.matrix[i][j])
			}
		}
		table += fmt.Sprintf("%4s", " ")
		table += "\n"
	}
	for i := 0; i < 40; i++ {
		table += "-"
	}
	table += "\n"
	SendMessage(GroupChatId, *event.Event.Message.ChatId, Text, table)
}
