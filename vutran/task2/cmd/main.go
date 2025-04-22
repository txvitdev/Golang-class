package main

import (
	"task2/internal/server"

	"go.uber.org/fx"
)

func main() {
	fx.New(
		server.Module,
		fx.Invoke(server.RunServer),
	).Run()
}
