package handler

import (
	"net/http"

	"github.com/RaffDv/goopportunities/internal/schema"
	"github.com/gin-gonic/gin"
)

func DeleteOpeningHandler(c *gin.Context) {
	id := c.Query("id")

	if id == "" {
		sendError(c, http.StatusBadRequest, errParamIsRequired("id", "queryParam").Error())
		return
	}

	opening := schema.Opening{}
	if err := db.First(&opening, id).Error; err != nil {
		sendError(c, http.StatusNotFound, err.Error())
		return
	}

	if err := db.Delete(&opening).Error; err != nil {
		sendError(c, http.StatusInternalServerError, err.Error())
		return
	}

	sendSuccess(c, http.StatusNoContent, "delete-opening", opening)

}
