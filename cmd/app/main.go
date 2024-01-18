package main

import (
	"log/slog"

	"github.com/murasame29/docker-bussiness-card/cmd/config"
	"github.com/murasame29/docker-bussiness-card/internal/router"
	"github.com/murasame29/docker-bussiness-card/internal/server"
)

func init() {
	config.LoadEnv()
}

func main() {
	router := router.NewRouter()
	server.New(router, slog.Default()).RunWithGracefulShutdown()
}
