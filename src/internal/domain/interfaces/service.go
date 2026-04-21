package interfaces

import "tictactoe/internal/domain/models"

type GameService interface {
	NextMove(game models.Game) (models.Game, error)

	ValidateBoard(current models.Game, updated models.Game) error

	IsGameOver(game models.Game) bool
}
