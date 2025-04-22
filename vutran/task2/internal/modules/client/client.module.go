package modules

import (
	handlers "task2/internal/handlers/client"
	repositories "task2/internal/repositories/client"
	services "task2/internal/services/client"

	"go.uber.org/fx"
)

var ClientModule = fx.Options(
	fx.Provide(repositories.NewClientRepository),
	fx.Provide(services.NewClientService),
	fx.Provide(handlers.NewClientHandler),
)
