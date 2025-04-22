package server

import (
	"context"
	authModule "task2/internal/modules/auth"
	clientModule "task2/internal/modules/client"
	configModule "task2/internal/modules/config"
	databaseModule "task2/internal/modules/database"
	roleModule "task2/internal/modules/role"
	"task2/internal/routes"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

var Module = fx.Options(
	clientModule.ClientModule,
	authModule.AuthModule,
	configModule.ConfigModule,
	databaseModule.DatabaseModule,
	roleModule.RoleModule,
	fx.Provide(routes.NewRouter),
)

func RunServer(lc fx.Lifecycle, router *gin.Engine) {
	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			go router.Run()
			return nil
		},
		OnStop: func(context.Context) error {
			return nil
		},
	})
}
