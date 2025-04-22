package modules

import (
	"task2/internal/config"

	"go.uber.org/fx"
)

var ConfigModule = fx.Options(
	fx.Provide(config.SetupConfig),
)
