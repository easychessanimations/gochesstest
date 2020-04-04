// https://medium.com/rungo/anatomy-of-modules-in-go-c8274d215c16

package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"

	"github.com/easychessanimations/gochess/board"
)

const WASM = false

func main() {
	board.About()

	b := board.Board{}

	b.Init(board.VARIANT_STANDARD)

	b.Reset()

	//b.SetFromFen("")

	if WASM {
		b.Perf(3)
	}

	reader := bufio.NewReader(os.Stdin)

	ok := true

	for ok {
		fmt.Println("")

		b.Print()

		lms := b.LegalMovesForAllPieces()

		mb := b.MovesSortedBySan(lms)

		fmt.Printf("Legal moves ( %d ) ", len(lms))

		for i, mbi := range mb {
			fmt.Printf("%d. %s [ %s ] ", i+1, mbi.San, mbi.Algeb)
		}

		fmt.Print("\n\n> ")
		text, _ := reader.ReadString('\n')

		text = strings.Trim(text, "\r\n")

		i, err := strconv.ParseInt(text, 10, 32)

		if err == nil {
			move := mb[i-1].Move

			b.Push(move)
		} else {
			if text == "perf" {
				b.Perf(3)
			} else if text == "d" {
				b.Pop()
			} else if text == "" {
				randIndex := rand.Intn(len(mb))

				move := mb[randIndex-1].Move

				b.Push(move)
			} else if text != "" {
				for _, mbi := range mb {
					if (mbi.San == text) || (mbi.Algeb == text) {
						move := mbi.Move

						b.Push(move)
					}
				}
			}
		}

		ok = text != "x"

		if WASM {
			ok = false
		}
	}
}
