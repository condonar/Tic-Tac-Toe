package di

import (
	"context"
	"net/http"
	"tictactoe/internal/datasource/repository"
	"tictactoe/internal/datasource/storage"
	"tictactoe/internal/domain/interfaces"
	"tictactoe/internal/domain/service"
	"tictactoe/internal/web/handler"

	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		storage.NewGameStorage,
		repository.NewGameRepository,
		asGameRepository,
		service.NewGameService,
		asGameService,
		handler.NewGameHandler,
		newMux,
		newHTTPServer,
	),
	fx.Invoke(registerHTTPServer),
)

func asGameRepository(r *repository.GameRepository) interfaces.GameRepository {
	return r
}

func asGameService(s *service.GameService) interfaces.GameService {
	return s
}

func newMux(h *handler.GameHandler) *http.ServeMux {
	mux := http.NewServeMux()
	h.RegisterRoutes(mux)
	return mux
}

func newHTTPServer(mux *http.ServeMux) *http.Server {
	return &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
}

func registerHTTPServer(lc fx.Lifecycle, server *http.Server) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				_ = server.ListenAndServe()
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return server.Shutdown(ctx)
		},
	})
}
