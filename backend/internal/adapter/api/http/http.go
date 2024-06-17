package http

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/rasouliali1379/terrax/config"
	"github.com/rasouliali1379/terrax/internal/core/ports"
	"github.com/rasouliali1379/terrax/internal/pkg/ginper"
	"go.uber.org/zap"
	"log/slog"
)

type Server struct {
	engine     *gin.Engine
	appService ports.AppService
}

func New(appService ports.AppService) Server {

	binding.Validator = NewValidator()

	if config.C().Env == "prod" {
		gin.SetMode(gin.ReleaseMode)
	}

	return Server{
		engine:     gin.New(),
		appService: appService,
	}
}

func (s Server) Serve() error {
	zap.S().Info("Running http server")

	s.engine.Use(gin.CustomRecoveryWithWriter(nil, ginper.PanicHandler))
	s.engine.Use(corsHandler())
	s.addRoutes()
	slog.Info(fmt.Sprintf("Http server is running on :%s", config.C().Http.Port))
	if err := s.engine.Run(":" + config.C().Http.Port); err != nil {
		zap.S().Info(err)
	}
	return nil
}

func corsHandler() gin.HandlerFunc {
	cfg := cors.DefaultConfig()
	cfg.AllowAllOrigins = true
	cfg.AllowCredentials = true
	cfg.AddAllowHeaders("Authorization")
	return cors.New(cfg)
}
