package repository

import (
	"fmt"
	"tictactoe/internal/datasource/mapper"
	"tictactoe/internal/datasource/storage"
	"tictactoe/internal/domain/models"

	"github.com/google/uuid"
)

type GameRepository struct {
	storage *storage.GameStorage
}

func NewGameRepository(storage *storage.GameStorage) *GameRepository {
	return &GameRepository{storage: storage}
}

func (r *GameRepository) Save(game models.Game) error {
	dsGame := mapper.ToDatasourceGame(game)
	r.storage.Save(dsGame)
	return nil
}

func (r *GameRepository) Get(id uuid.UUID) (models.Game, error) {
	dsGame, ok := r.storage.Get(id.String())
	if !ok {
		return models.Game{}, fmt.Errorf("game not found: %s", id.String())
	}

	game, err := mapper.ToDomainGame(dsGame)
	if err != nil {
		return models.Game{}, err
	}

	return game, nil
}
