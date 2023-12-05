package handler

import (
	"net/http"

	"github.com/RaffDv/goopportunities/internal/schema"
	"github.com/gin-gonic/gin"
)

func UpdateOpeningHandler(c *gin.Context) {
	body := UpdateOpeningRequest{}
	id := c.Query("id")

	c.BindJSON(&body)
	if err := body.Validate(); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		sendError(c, http.StatusUnprocessableEntity, err.Error())
	}

	if id == "" {
		logger.Error("ID not provided")
		sendError(c, http.StatusBadRequest, errParamIsRequired("id", "queryParam").Error())
		return
	}

	opening := schema.Opening{}
	if err := db.First(&opening, id).Error; err != nil {
		logger.Errorf("not found opening with id: %s", id)
		sendError(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := db.Model(&opening).Updates(&body).Error; err != nil {
		logger.Errorf("Error updating data in database: %v", err.Error())
		sendError(c, http.StatusUnprocessableEntity, "error updating opening")
		return
	}

	sendSuccess(c, http.StatusNoContent, "update-opening", opening)
}
