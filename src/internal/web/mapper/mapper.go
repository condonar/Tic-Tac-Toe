package mapper

import (
	"fmt"

	domainmodels "tictactoe/internal/domain/models"
	webmodels "tictactoe/internal/web/models"

	"github.com/google/uuid"
)

func ToWebGame(game domainmodels.Game) webmodels.Game {
	return webmodels.Game{
		ID: game.ID.String(),
		Board: webmodels.Board{
			Cells: copyCells(game.Board.Cells),
		},
	}
}

func ToDomainGame(game webmodels.Game) (domainmodels.Game, error) {
	id, err := uuid.Parse(game.ID)
	if err != nil {
		return domainmodels.Game{}, fmt.Errorf("invalid game ID: %w", err)
	}

	return domainmodels.Game{
		ID: id,
		Board: domainmodels.Board{
			Cells: copyCells(game.Board.Cells),
		},
	}, nil
}

func copyCells(src [][]int) [][]int {
	if src == nil {
		return nil
	}
	dst := make([][]int, len(src))
	for i := range src {
		dst[i] = make([]int, len(src[i]))
		copy(dst[i], src[i])
	}
	return dst
}
