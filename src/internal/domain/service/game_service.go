package service

import (
	"errors"

	"tictactoe/internal/domain/interfaces"
	"tictactoe/internal/domain/models"
)

type GameService struct {
	repo interfaces.GameRepository
}

func NewGameService(repo interfaces.GameRepository) *GameService {
	return &GameService{repo: repo}
}

var _ interfaces.GameService = (*GameService)(nil)

func (s *GameService) NextMove(game models.Game) (models.Game, error) {
	if s.IsGameOver(game) {
		return models.Game{}, errors.New("NextMove: game already over")
	}

	board := copyBoard(game.Board.Cells)
	bestScore := -2
	bestRow, bestCol := -1, -1

	for r := 0; r < models.BoardSize; r++ {
		for c := 0; c < models.BoardSize; c++ {
			if board[r][c] != models.Empty {
				continue
			}
			board[r][c] = models.PlayerO
			score := minimax(board, false)
			board[r][c] = models.Empty

			if score > bestScore {
				bestScore = score
				bestRow, bestCol = r, c
			}
		}
	}
	if bestRow < 0 {
		return models.Game{}, errors.New("NextMove: no valid moves")
	}

	board[bestRow][bestCol] = models.PlayerO
	updated := models.Game{
		ID:    game.ID,
		Board: models.Board{Cells: board},
	}

	if err := s.repo.Save(updated); err != nil {
		return models.Game{}, err
	}

	return updated, nil
}

func (s *GameService) ValidateBoard(current models.Game, updated models.Game) error {
	if current.ID != updated.ID {
		return errors.New("Validateboard: id mismatch")
	}

	cur := current.Board.Cells
	upd := updated.Board.Cells

	if !isValidShape(upd) {
		return errors.New("Validateboard: invalid board size")
	}

	if s.IsGameOver(current) {
		return errors.New("Validateboard: game already over")
	}

	newMoves := 0

	for r := 0; r < models.BoardSize; r++ {
		for c := 0; c < models.BoardSize; c++ {
			cell := upd[r][c]

			// допустимые значения: X, O, Empty
			if cell != models.Empty && cell != models.PlayerX && cell != models.PlayerO {
				return errors.New("Validateboard: invalid cell value")
			}

			if cur[r][c] != models.Empty {
				// Старая клетка не должна меняться
				if upd[r][c] != cur[r][c] {
					return errors.New("Validateboard: board was altered")
				}
			} else if upd[r][c] != models.Empty {
				// новая заполненная клетка - ход игрока
				newMoves++
				if upd[r][c] != models.PlayerX {
					return errors.New("Validateboard: expected X move")
				}
			}
		}
	}

	if newMoves != 1 {
		return errors.New("Validateboard: expected exactly one new move")
	}

	return nil
}

func (s *GameService) IsGameOver(game models.Game) bool {
	board := game.Board.Cells

	if hasWinner(board) {
		return true
	}

	for r := 0; r < models.BoardSize; r++ {
		for c := 0; c < models.BoardSize; c++ {
			if board[r][c] == models.Empty {
				return false
			}
		}
	}
	return true
}
