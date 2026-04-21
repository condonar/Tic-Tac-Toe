package handler

import (
	"encoding/json"
	"net/http"

	"tictactoe/internal/domain/interfaces"
	domainmodels "tictactoe/internal/domain/models"
	"tictactoe/internal/web/mapper"
	webmodels "tictactoe/internal/web/models"

	"github.com/google/uuid"
)

type GameHandler struct {
	service interfaces.GameService
	repo    interfaces.GameRepository
}

func NewGameHandler(service interfaces.GameService, repo interfaces.GameRepository) *GameHandler {
	return &GameHandler{service: service, repo: repo}
}

func (h *GameHandler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /game/{id}", h.HandleMove)
}

func (h *GameHandler) HandleMove(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid game id in path")
		return
	}

	var req webmodels.Game
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid json body")
		return
	}
	req.ID = id.String()

	updated, err := mapper.ToDomainGame(req)
	if err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	current, err := h.repo.Get(id)
	if err != nil {
		current = domainmodels.NewEmptyGame(id)
	}

	if err := h.service.ValidateBoard(current, updated); err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}
	if h.service.IsGameOver(updated) {
		if err := h.repo.Save(updated); err != nil {
			writeError(w, http.StatusInternalServerError, err.Error())
			return
		}
		writeJSON(w, http.StatusOK, mapper.ToWebGame(updated))
		return
	}

	result, err := h.service.NextMove(updated)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	writeJSON(w, http.StatusOK, mapper.ToWebGame(result))
}
