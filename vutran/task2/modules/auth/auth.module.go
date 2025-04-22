package modules

import (
	controllers "task2/controllers/auth"
	repositories "task2/repositories/user"
	services "task2/services/auth"

	"go.uber.org/fx"
)

var AuthModule = fx.Options(
	fx.Provide(repositories.NewUserRepository),
	fx.Provide(services.NewAuthService),
	fx.Provide(controllers.NewAuthController),
)