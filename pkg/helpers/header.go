package helpers

import (
	"prodtrack-api/pkg/errs"

	"github.com/gin-gonic/gin"
)

func CheckContentType(ctx *gin.Context) errs.ErrorResponse {
	if ctx.GetHeader("Content-Type") != "application/json" {
		return errs.NewUnsupportedMediaTypeError("Unsupported content-type. Only 'application/json' is supported.")
	}

	return nil
}
