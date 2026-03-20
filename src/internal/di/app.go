package di

import (
	"Project03-Go_Bootcamp/internal/datasource"
	"Project03-Go_Bootcamp/internal/services/minimax"
	"Project03-Go_Bootcamp/internal/services/register"
	"Project03-Go_Bootcamp/internal/web/game_post"
	"Project03-Go_Bootcamp/internal/web/get_html_get"
	"Project03-Go_Bootcamp/internal/web/register_game_id_get"
	"context"
	"log"
	"net/http"

	"go.uber.org/fx"
)

func NewMux() *http.ServeMux {
	return http.NewServeMux()
}

func RegisterRoutes(h *game_post.Handler, h2 *get_html_get.Handler, h3 *register_game_id_get.Handler, mux *http.ServeMux) {
	h.Register(mux)
	h2.Register(mux)
	h3.Register(mux)
}

func RunHTTPServer(lc fx.Lifecycle, mux *http.ServeMux) {
	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				log.Println("HTTP listen :8080")
				if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
					log.Printf("http server error: %v", err)
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			log.Println("HTTP shutdown")
			return server.Shutdown(ctx)
		},
	})
}

var Module = fx.Options(
	fx.Provide(
		fx.Annotate(
			datasource.New,
			fx.As(new(minimax.Repository)),
			fx.As(new(register.Repository)),
		),
	),

	fx.Provide(
		fx.Annotate(
			minimax.New,
			fx.As(new(game_post.Service)),
		),
	),

	fx.Provide(
		fx.Annotate(
			register.New,
			fx.As(new(register_game_id_get.Service)),
		),
	),

	fx.Provide(NewMux),
	fx.Provide(get_html_get.New),
	fx.Provide(register_game_id_get.New),
	fx.Provide(game_post.New),

	fx.Invoke(RegisterRoutes),
	fx.Invoke(RunHTTPServer),
)
