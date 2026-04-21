package models

const (
	Empty   = 0
	PlayerX = 1
	PlayerO = 2
)

const BoardSize = 3

type Board struct {
	Cells [][]int
}
