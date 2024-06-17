package http

import (
	"github.com/gin-gonic/gin"
	"github.com/rasouliali1379/terrax/config"
)

func (s Server) addRoutes() *gin.Engine {

	router := s.engine.RouterGroup.Group("").Use(s.authMiddleware(config.C().Telegram.BotToken))
	router.GET("/", s.get)
	router.POST("/", s.set)

	return s.engine
}
