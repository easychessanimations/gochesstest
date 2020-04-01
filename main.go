// https://medium.com/rungo/anatomy-of-modules-in-go-c8274d215c16

package main

import (
	"github.com/easychessanimations/gochess/board"
)

func main() {
	board.About()

	b := board.Board{}

	b.Init(board.VARIANT_STANDARD)

	b.Print()
}
