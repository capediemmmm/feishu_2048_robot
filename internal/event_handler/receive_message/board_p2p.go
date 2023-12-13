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

func p2pNewGame(event *larkim.P2MessageReceiveV1) {

	// var table string
	/*
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
	*/
	// table = "2048\naction!"
	g = New()
	g.AddElement()
	g.AddElement()
	g.Display(event)
	// SendMessage(UserOpenId, *event.Event.Sender.SenderId.OpenId, Text, table)
}

func New() Board {
	matrix := make([][]int, 0)
	for i := 0; i < _rows; i++ {
		matrix = append(matrix, make([]int, _cols))
	}
	return &board{
		matrix: matrix,
	}
}

func (b *board) AddElement() {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	val := r1.Int() % 100 // 0-99
	if val <= 50 {
		val = 2
	} else {
		val = 4
	}

	empty := 0
	for i := 0; i < _rows; i++ {
		for j := 0; j < _cols; j++ {
			if b.matrix[i][j] == 0 {
				empty++
			}
		}
	}
	elementCount := r1.Int()%empty + 1
	index := 0

	for i := 0; i < _rows; i++ {
		for j := 0; j < _cols; j++ {
			if b.matrix[i][j] == 0 {
				index++
				if index == elementCount {
					b.newRow = i
					b.newCol = j
					b.matrix[i][j] = val
					return
				}
			}
		}
	}
	return
}

func (b *board) Display(event *larkim.P2MessageReceiveV1) {
	var table string
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
	SendMessage(UserOpenId, *event.Event.Sender.SenderId.OpenId, Text, table)
}

func (b *board) Move(event *larkim.P2MessageReceiveV1, content string) {
	switch content {
	case "左":
		b.moveLeft()
	case "右":
		b.moveRight()
	case "上":
		b.moveUp()
	case "下":
		b.moveDown()
	}
}

func (b *board) moveLeft() {
	for i := 0; i < _rows; i++ {
		old := b.matrix[i]
		b.matrix[i] = movedRow(old)
	}
}

func (b *board) moveUp() {
	b.reverseRows()
	b.moveDown()
	b.reverseRows()
}

func (b *board) moveDown() {
	b.transpose()
	b.moveLeft()
	b.transpose()
	b.transpose()
	b.transpose()
}

func (b *board) moveRight() {
	b.reverse()
	b.moveLeft()
	b.reverse()
}

func movedRow(elems []int) []int {
	nonEmpty := make([]int, 0)
	for i := 0; i < _cols; i++ {
		if elems[i] != 0 {
			nonEmpty = append(nonEmpty, elems[i])
		}
	}
	remaining := _cols - len(nonEmpty)
	for i := 0; i < remaining; i++ {
		nonEmpty = append(nonEmpty, 0)
	}
	return mergeElements(nonEmpty)
}

// reverse simply reverses each row of the board
func (b *board) reverse() {
	for i := 0; i < _rows; i++ {
		b.matrix[i] = reverseRow(b.matrix[i])
	}
}

// transpose rotates a list
// row becomes _cols
// [ 1 2 ]
// [ 3 4 ] becomes
//
// [ 3 1 ]
// [ 4 2 ]
// see test for more clarity
func (b *board) transpose() {
	ans := make([][]int, 0)
	for i := 0; i < _rows; i++ {
		ans = append(ans, make([]int, _cols))
	}
	for i := 0; i < _rows; i++ {
		for j := 0; j < _cols; j++ {
			ans[i][j] = b.matrix[_cols-j-1][i]
		}
	}
	b.matrix = ans
}

// reverseRows reverses the order of lists
// [1 2]
// [3 4] becomes
//
// [3 4]
// [1 2]
func (b *board) reverseRows() {
	ans := make([][]int, 0)
	for i := 0; i < _rows; i++ {
		ans = append(ans, make([]int, _cols))
	}
	for i := 0; i < _rows; i++ {
		for j := 0; j < _cols; j++ {
			ans[_rows-i-1][j] = b.matrix[i][j]
		}
	}
	b.matrix = ans
}

// reverseRow reverses a row
func reverseRow(arr []int) []int {
	ans := make([]int, 0)
	for i := len(arr) - 1; i >= 0; i-- {
		ans = append(ans, arr[i])
	}
	return ans
}

// mergeElements when a row is moved to left, it merges the element which can
// see tests for more clarity
func mergeElements(arr []int) []int {
	newArr := make([]int, len(arr))
	newArr[0] = arr[0]
	index := 0
	for i := 1; i < len(arr); i++ {
		if arr[i] == newArr[index] {
			newArr[index] += arr[i]
		} else {
			index++
			newArr[index] = arr[i]
		}
	}
	return newArr
}

func (b *board) IsOver() bool {
	empty := 0
	for i := 0; i < _rows; i++ {
		for j := 0; j < _cols; j++ {
			if b.matrix[i][j] == 0 {
				empty++
			}
		}
	}
	return empty == 0
}

func (b *board) CountScore() (int, int) {
	max := 0
	sum := 0
	for i := 0; i < _rows; i++ {
		for j := 0; j < _cols; j++ {
			sum += b.matrix[i][j]
			if b.matrix[i][j] > max {
				max = b.matrix[i][j]
			}
			if b.matrix[i][j] == 2048 {
				b.success = true
			}
		}
	}
	return sum, max
}

func (b *board) IsSuccess() bool {
	return b.success
}
