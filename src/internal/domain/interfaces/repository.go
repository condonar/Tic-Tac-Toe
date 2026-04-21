package interfaces

import (
	"tictactoe/internal/domain/models"

	"github.com/google/uuid"
)

type GameRepository interface {
	Save(game models.Game) error
	Get(id uuid.UUID) (models.Game, error)
}
