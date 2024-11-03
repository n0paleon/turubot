package main

import (
	"go.uber.org/fx"
	"turubot/infra/config"
	"turubot/infra/logger"
	"turubot/infra/workerpool"
	"turubot/internal/adapters/waifupics"
	"turubot/internal/application/bot"
	"turubot/internal/domain"
)

// bot entrypoint
func main() {
	cfg := config.LoadConfig("config.yaml")
	consoleLogger := logger.NewConsoleLogger(cfg)
	_ = workerpool.InitializePool(cfg.PoolSize, consoleLogger)

	app := fx.New(
		fx.Provide(func() *config.Config { return cfg }),
		fx.Provide(func() domain.Logger { return consoleLogger }),
		fx.Provide(bot.NewBotApp),
		waifupics.Module,
		fx.Invoke(func(b *bot.TelebotApp) {
			b.HandleIncomingMessages()
		}),
	)
	// run application
	app.Run()
}
