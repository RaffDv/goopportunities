package handler

import (
	"fmt"
	"net/http"

	"github.com/RaffDv/goopportunities/internal/schema"
	"github.com/gin-gonic/gin"
)

func FindOpeningHandler(c *gin.Context) {
	id := c.Query("id")

	if id == "" {
		sendError(c, http.StatusBadRequest, errParamIsRequired("id", "queryParam").Error())
		return
	}

	opening := schema.Opening{}

	db.First(&opening, id)
	if opening.ID == 0 {
		sendError(c, http.StatusNotFound, fmt.Sprintf("opening with id : %s not found", id))
		return
	}

	sendSuccess(c, http.StatusOK, "find-opening", opening)
}
