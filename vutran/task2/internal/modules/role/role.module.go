package modules

import (
	handlers "task2/internal/handlers/role"
	repositories "task2/internal/repositories/role"
	services "task2/internal/services/role"

	"go.uber.org/fx"
)

var RoleModule = fx.Options(
	fx.Provide(repositories.NewRoleRepository),
	fx.Provide(services.NewRoleService),
	fx.Provide(handlers.NewRoleHandler),
)
