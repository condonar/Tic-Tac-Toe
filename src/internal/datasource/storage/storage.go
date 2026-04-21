package storage

import (
	"sync"
	"tictactoe/internal/datasource/models"
)

type GameStorage struct {
	games sync.Map
}

func NewGameStorage() *GameStorage {
	return &GameStorage{}
}

func (s *GameStorage) Save(game models.Game) {
	s.games.Store(game.ID, game)
}

func (s *GameStorage) Get(id string) (models.Game, bool) {
	value, ok := s.games.Load(id)
	if !ok {
		return models.Game{}, false
	}
	return value.(models.Game), true
}
