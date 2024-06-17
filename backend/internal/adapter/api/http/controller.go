package http

import (
	"github.com/gin-gonic/gin"
	"github.com/rasouliali1379/terrax/internal/pkg/ginper"
)

func (s Server) set(c *gin.Context) {
	initData, ok := ctxInitData(c.Request.Context())
	if !ok {
		ginper.Forbidden(c, "init data not found")
		return
	}

	ginper.SuccessWithBody(c, initData)
}

func (s Server) get(c *gin.Context) {
	initData, ok := ctxInitData(c.Request.Context())
	if !ok {
		ginper.Forbidden(c, "init data not found")
		return
	}

	res, err := s.appService.GetUserData(c, mapInitDataToUserModel(initData))
	if err != nil {
		handleError(c, err)
		return
	}

	ginper.SuccessWithBody(c, res)
}
