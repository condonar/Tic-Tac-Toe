package service

import "tictactoe/internal/domain/models"

func copyBoard(src [][]int) [][]int {
	dst := make([][]int, len(src))
	for i := range src {
		dst[i] = make([]int, len(src[i]))
		copy(dst[i], src[i])
	}
	return dst
}

func minimax(board [][]int, isMaximizing bool) int {
	w := winner(board)
	if w == models.PlayerO {
		return 1
	}
	if w == models.PlayerX {
		return -1
	}
	if !hasEmptyCell(board) {
		return 0
	}

	if isMaximizing {
		best := -2
		for r := 0; r < models.BoardSize; r++ {
			for c := 0; c < models.BoardSize; c++ {
				if board[r][c] != models.Empty {
					continue
				}
				board[r][c] = models.PlayerO
				score := minimax(board, false)
				board[r][c] = models.Empty
				if score > best {
					best = score
				}
			}
		}
		return best
	}

	best := 2
	for r := 0; r < models.BoardSize; r++ {
		for c := 0; c < models.BoardSize; c++ {
			if board[r][c] != models.Empty {
				continue
			}
			board[r][c] = models.PlayerX
			score := minimax(board, true)
			board[r][c] = models.Empty
			if score < best {
				best = score
			}
		}
	}
	return best
}

func hasEmptyCell(board [][]int) bool {
	for r := 0; r < models.BoardSize; r++ {
		for c := 0; c < models.BoardSize; c++ {
			if board[r][c] == models.Empty {
				return true
			}
		}
	}
	return false
}

func isValidShape(board [][]int) bool {
	if len(board) != models.BoardSize {
		return false
	}
	for r := 0; r < models.BoardSize; r++ {
		if len(board[r]) != models.BoardSize {
			return false
		}
	}
	return true
}

func winner(board [][]int) int {
	for r := 0; r < models.BoardSize; r++ {
		if board[r][0] != models.Empty &&
			board[r][0] == board[r][1] &&
			board[r][1] == board[r][2] {
			return board[r][0]
		}
	}
	for c := 0; c < models.BoardSize; c++ {
		if board[0][c] != models.Empty &&
			board[0][c] == board[1][c] &&
			board[1][c] == board[2][c] {
			return board[0][c]
		}
	}
	if board[0][0] != models.Empty &&
		board[0][0] == board[1][1] &&
		board[1][1] == board[2][2] {
		return board[0][0]
	}
	if board[0][2] != models.Empty &&
		board[0][2] == board[1][1] &&
		board[1][1] == board[2][0] {
		return board[0][2]
	}
	return models.Empty
}

func hasWinner(board [][]int) bool {
	return winner(board) != models.Empty
}
