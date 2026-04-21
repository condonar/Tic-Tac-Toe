package mapper

import (
	"fmt"

	dsmodels "tictactoe/internal/datasource/models"
	domainmodels "tictactoe/internal/domain/models"

	"github.com/google/uuid"
)

func ToDatasourceGame(game domainmodels.Game) dsmodels.Game {
	return dsmodels.Game{
		ID: game.ID.String(),
		Board: dsmodels.Board{
			Cells: copyCells(game.Board.Cells),
		},
	}
}

func ToDomainGame(game dsmodels.Game) (domainmodels.Game, error) {
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
