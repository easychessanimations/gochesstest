// https://medium.com/rungo/anatomy-of-modules-in-go-c8274d215c16

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/easychessanimations/gochess/board"
)

func main() {
	board.About()

	b := board.Board{}

	b.Init(board.VARIANT_EIGHTPIECE)

	b.Reset()

	reader := bufio.NewReader(os.Stdin)

	ok := true

	for ok {
		fmt.Println("")

		b.Print()

		pslms := b.PslmsForAllPiecesOfColor(b.Pos.Turn)

		mb := b.MovesSortedBySan(pslms)

		for i, mbi := range mb {
			fmt.Printf("%d. %s ", i+1, mbi.Str)
		}

		fmt.Print("\n\n> ")
		text, _ := reader.ReadString('\n')

		text = strings.Trim(text, "\r\n")

		i, err := strconv.ParseInt(text, 10, 32)

		if err == nil {
			move := mb[i-1].Move

			b.Push(move)
		} else {
			if text == "d" {
				b.Pop()
			} else {
				for _, mbi := range mb {
					if mbi.Str == text {
						move := mbi.Move

						b.Push(move)
					}
				}
			}
		}

		ok = text != "x"
	}
}
