package sgin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type httpError struct {
	code    int
	message string
	err     error
}

func (e *httpError) Error() string {
	return e.message
}

func abortError(ctx *gin.Context, err error) {
	switch v := err.(type) {
	case *httpError:
		ctx.AbortWithStatus(v.code)
	default:
		ctx.AbortWithStatus(http.StatusInternalServerError)
	}
}
