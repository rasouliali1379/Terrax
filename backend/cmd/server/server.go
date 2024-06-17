package server

import (
	"github.com/rasouliali1379/terrax/config"
	"github.com/rasouliali1379/terrax/internal/adapter/api/http"
	"github.com/rasouliali1379/terrax/internal/adapter/infra/mongodb"
	"github.com/rasouliali1379/terrax/internal/pkg/logger"
	"github.com/rasouliali1379/terrax/internal/pkg/terminator"
	"github.com/spf13/cobra"
)

func Start(_ *cobra.Command, args []string) {
	config.Init()
	logger.Init()

	_, mongoShutdown := mongodb.Init()

	terminator.NewTerminator([]terminator.Func{
		mongoShutdown,
	}).Listen()

	if err := http.New().Serve(); err != nil {
		panic(err)
	}
}
