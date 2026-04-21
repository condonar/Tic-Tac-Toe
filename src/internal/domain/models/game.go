package models

import "github.com/google/uuid"

type Game struct {
	ID    uuid.UUID
	Board Board
}

func NewEmptyGame(id uuid.UUID) Game {
	cells := make([][]int, BoardSize)
	for i := range cells {
		cells[i] = make([]int, BoardSize)
	}
	return Game{ID: id, Board: Board{Cells: cells}}
}
