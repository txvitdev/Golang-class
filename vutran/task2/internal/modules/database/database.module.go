package modules

import (
	"task2/internal/database"

	"go.uber.org/fx"
)

var DatabaseModule = fx.Options(
	fx.Provide(database.NewDb),
)
