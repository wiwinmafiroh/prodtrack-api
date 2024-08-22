package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Welcome godoc
// @ID welcome-message
// @Summary Welcome message
// @Description This endpoint provides a simple welcome message. It doesn't require authentication.
// @Tags General
// @Produce text/html
// @Success 200 {string} string
// @Router / [get]
func Welcome(ctx *gin.Context) {
	ctx.Header("Content-Type", "text/html; charset=utf-8")
	ctx.String(http.StatusOK, "Welcome to the Prodtrack API. Manage your products here!")
}
