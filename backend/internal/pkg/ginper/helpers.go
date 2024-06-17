package ginper

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
	"log/slog"
	"net/http"
	"runtime/debug"
	"strings"
)

func ServerError(c *gin.Context, e error) {
	if err := errors.Unwrap(e); err != nil {
		if errors.Is(err, context.Canceled) ||
			errors.Is(err, context.DeadlineExceeded) ||
			strings.Contains(e.Error(), "context deadline exceeded") {
			c.AbortWithStatus(http.StatusRequestTimeout)
			return
		}
	}

	slog.Error("internal error", e, slog.String("stack_trace", string(debug.Stack())))

	c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "internal server error occurred"})
}

func PanicHandler(c *gin.Context, e any) {
	err, ok := e.(error)
	if !ok {
		err = errors.New(fmt.Sprintf("%s", e))
	}

	if err := errors.Unwrap(err); err != nil {
		if errors.Is(err, context.Canceled) || errors.Is(err, context.DeadlineExceeded) {
			return
		}
	}

	slog.Error("internal error", err, slog.String("stack_trace", string(debug.Stack())))
	c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "internal server error occurred"})
}

func BadRequest(c *gin.Context, msg string) {
	c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": msg})
}

func Success(c *gin.Context, msg string) {
	c.AbortWithStatusJSON(http.StatusOK, gin.H{"message": msg})
}

func Created(c *gin.Context, msg string) {
	c.AbortWithStatusJSON(http.StatusCreated, gin.H{"message": msg})
}

func CreatedWithBody(c *gin.Context, body any) {
	c.AbortWithStatusJSON(http.StatusCreated, body)
}

func SuccessWithBody(c *gin.Context, body any) {
	c.AbortWithStatusJSON(http.StatusOK, body)
}

func NotFound(c *gin.Context, msg string) {
	c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": msg})
}

func NoContent(c *gin.Context) {
	c.AbortWithStatus(http.StatusNoContent)
}

func NoContentWithError(c *gin.Context, e error) {
	slog.Error("service unavailable", e)
	c.AbortWithStatus(http.StatusNoContent)
}

func Forbidden(c *gin.Context, e string) {
	c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": e})
}

func Unauthorized(c *gin.Context, e string) {
	c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": e})
}

func Csv(c *gin.Context, content []byte, fileName string) {
	c.Status(http.StatusOK)
	c.Header("Content-Type", "text/csv")
	c.Header("Content-Disposition", fmt.Sprintf("attachment;filename=%s.csv", fileName))
	_, err := c.Writer.Write(content)
	if err != nil {
		ServerError(c, err)
		return
	}
}
