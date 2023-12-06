package receiveMessage

import (
	"fmt"
	"math/rand"
	"time"

	larkim "github.com/larksuite/oapi-sdk-go/v3/service/im/v1"
)

func p2pHelpMenu(event *larkim.P2MessageReceiveV1) {
	SendMessage(UserOpenId, *event.Event.Sender.SenderId.OpenId, Text, "this is a p2p test string")
}

func p2pMove(event *larkim.P2MessageReceiveV1) {
	var rows int
	var cols int
	rows = 4
	cols = 4
	table := make([][]string, rows)
	cellIndex := 0
	for i := 0; i < rows; i++ {
		table[i] = make([]string, cols)
		for j := 0; j < cols; j++ {
			// 解析单元格内容
			for cellIndex < len(tableHTML) && tableHTML[cellIndex] != '<' {
				cellIndex++
			}
			start := cellIndex
			for cellIndex < len(tableHTML) && tableHTML[cellIndex] != '>' {
				cellIndex++
			}
			end := cellIndex
			table[i][j] = tableHTML[start : end+1]
		}
	}

	// 将每一行整体左移
	for i := 0; i < rows; i++ {
		// 保存第一个单元格的内容
		firstCell := table[i][0]
		// 将每个单元格左移
		for j := 0; j < cols-1; j++ {
			table[i][j] = table[i][j+1]
		}
		// 将第一个单元格放到最后一个位置
		table[i][cols-1] = firstCell
	}

	// 将二维切片重新转换为HTML字符串
	movedTableHTML := "<table>"
	for i := 0; i < rows; i++ {
		movedTableHTML += "<tr>"
		for j := 0; j < cols; j++ {
			movedTableHTML += table[i][j]
		}
		movedTableHTML += "</tr>"
	}
	movedTableHTML += "</table>"
}

func p2pNewGame(event *larkim.P2MessageReceiveV1) {
	var table string
	rand.New(rand.NewSource(time.Now().UnixNano()))
	tableHTML := `<table border="4">`

	// Initialize an empty 4x4 table
	for i := 0; i < 4; i++ {
		tableHTML += "<tr>"
		for j := 0; j < 4; j++ {
			tableHTML += "<td>0</td>"
		}
		tableHTML += "</tr>"
	}

	// Randomly initialize two cells with the value 2
	for k := 0; k < 2; k++ {
		row := rand.Intn(4)
		col := rand.Intn(4)
		tableHTML = tableHTML[:((row)*4)+col] + "2" + tableHTML[(row)*4+col+1:]
	}

	tableHTML += `</table>`
	table = fmt.Sprintf("<div>%s</div>", tableHTML)
	SendMessage(UserOpenId, *event.Event.Sender.SenderId.OpenId, Text, table)
}
