package server

import (
	authModule "task2/modules/auth"
	configModule "task2/modules/config"

	"go.uber.org/fx"
)

var Module = fx.Options(
	authModule.AuthModule,
	configModule.ConfigModule,
)
