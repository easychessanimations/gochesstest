// https://medium.com/rungo/anatomy-of-modules-in-go-c8274d215c16

package main

import (
	"fmt"

	"github.com/easychessanimations/gochess/board"
)

func main() {
	board.About()

	b := board.Board{}

	b.Init(board.VARIANT_STANDARD)

	b.Reset()

	b.Print()

	pslms := b.PslmsForAllPiecesOfColor(board.BLACK)

	for i, mbi := range b.MovesSortedBySan(pslms) {
		fmt.Printf("%d. %s\n", i+1, mbi.Str)
	}
}
