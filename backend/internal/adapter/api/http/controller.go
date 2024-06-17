package http

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/rasouliali1379/terrax/internal/pkg/ginper"
	"log/slog"
)

func (s Server) set(c *gin.Context) {
	initData, ok := ctxInitData(c.Request.Context())
	if !ok {
		ginper.Forbidden(c, errors.New("init data not found"))
		return
	}

	slog.Info("", initData)
	ginper.SuccessWithBody(c, initData)
}

func (s Server) get(c *gin.Context) {
	initData, ok := ctxInitData(c.Request.Context())
	if !ok {
		ginper.Forbidden(c, errors.New("init data not found"))
		return
	}

	ginper.SuccessWithBody(c, initData)
}
