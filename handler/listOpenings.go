package handler

import (
	"net/http"

	"github.com/becardine/gopportunities/schemas"
	"github.com/gin-gonic/gin"
)

func ListOpeningsHandler(ctx *gin.Context) {
	openings := []schemas.Opening{}

	if err := db.Find(&openings).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "failed to fetch openings")
		return
	}

	sendSuccess(ctx, "list-openings", openings)
}
