package modules

import (
	handlers "task2/internal/handlers/auth"
	repositories "task2/internal/repositories/user"
	services "task2/internal/services/auth"
	jwtService "task2/internal/services/jwt"

	"go.uber.org/fx"
)

var AuthModule = fx.Options(
	fx.Provide(repositories.NewUserRepository),
	fx.Provide(jwtService.ProvideJWTMaker),
	fx.Provide(services.NewAuthService),
	fx.Provide(handlers.NewAuthHandler),
)
