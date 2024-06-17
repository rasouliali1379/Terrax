package http

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/rasouliali1379/terrax/internal/core/domain"
	"github.com/rasouliali1379/terrax/internal/pkg/errorium"
	"github.com/rasouliali1379/terrax/internal/pkg/ginper"
)

func handleError(c *gin.Context, err error) {
	var e *errorium.Error
	if errors.As(err, &e) {
		switch e.Code() {
		case domain.ErrorNotFound:
			ginper.NotFound(c, e.Error())
			return
		case domain.ErrorNoContent:
			ginper.NoContent(c)
			return
		case domain.ErrorBadRequest:
			ginper.BadRequest(c, e.Error())
			return
		case domain.ErrorUnauthorized:
			ginper.Unauthorized(c, e.Error())
			return
		case domain.ErrorForbidden:
			ginper.Forbidden(c, e.Error())
			return
		default:
			ginper.ServerError(c, err)
			return
		}
	}

	ginper.ServerError(c, err)
}
