package handler

import (
	"net/http"

	"github.com/RaffDv/goopportunities/internal/schema"
	"github.com/gin-gonic/gin"
)

func ListOpeningsHandler(c *gin.Context) {

	openings := []schema.Opening{}

	db.Find(&openings)

	sendSuccess(c, http.StatusOK, "list-openings", openings)
}
