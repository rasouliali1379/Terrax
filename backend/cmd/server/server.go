package server

import (
	"github.com/rasouliali1379/terrax/config"
	"github.com/rasouliali1379/terrax/internal/adapter/api/http"
	"github.com/rasouliali1379/terrax/internal/adapter/infra/mongodb"
	"github.com/rasouliali1379/terrax/internal/adapter/repository"
	"github.com/rasouliali1379/terrax/internal/core/service"
	"github.com/rasouliali1379/terrax/internal/pkg/logger"
	"github.com/rasouliali1379/terrax/internal/pkg/terminator"
	"github.com/spf13/cobra"
)

func Start(_ *cobra.Command, args []string) {
	config.Init()
	logger.Init()

	db, mongoShutdown := mongodb.Init()

	terminator.NewTerminator([]terminator.Func{
		mongoShutdown,
	}).Listen()

	profileRepo := repository.NewProfileRepo(db)

	appService := service.NewAppService(profileRepo)

	if err := http.New(appService).Serve(); err != nil {
		panic(err)
	}
}
