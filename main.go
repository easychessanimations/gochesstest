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

	b.SetPieceAtSquare(board.Square{0, 2}, board.Piece{Kind: board.Rook, Color: board.WHITE})

	b.Print()

	pslms := b.PslmsForVectorPieceAtSquare(board.Piece{Kind: board.Knight}, board.Square{1, 0})

	for i, pslm := range pslms {
		fmt.Println(i, b.MoveToSan(pslm))
	}
}
