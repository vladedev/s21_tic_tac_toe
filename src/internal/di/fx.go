package di

import (
	"tic-tac-toe/internal/application"
	"tic-tac-toe/internal/datastorage"
	"tic-tac-toe/internal/domain"
	"tic-tac-toe/internal/web"

	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		datastorage.NewGameStorage,    // создание хранилища, singleton
		datastorage.NewGameRepository, // создание репозитория для исп. хранилища игры,
		application.NewGameService,    // создание сервиса в котором обрабатывается ход
		web.NewHandler,                // обработка post, get запросов
	),
	fx.Invoke(func(domain.GameRepository) {}), // гарантируем инициализацию
)
