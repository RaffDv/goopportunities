package handler

import (
	"net/http"

	"github.com/RaffDv/goopportunities/internal/schema"
	"github.com/gin-gonic/gin"
)

func UpdateOpeningHandler(c *gin.Context) {
	body := CreateOpeningRequest{}
	id := c.Query("id")

	if id == "" {
		sendError(c, http.StatusBadRequest, errParamIsRequired("id", "queryParam").Error())
		return
	}

	c.BindJSON(&body)

	opening := schema.Opening{}
	db.First(&opening, id)

	if err := db.Model(&opening).Updates(&body).Error; err != nil {
		logger.Errorf("Error creating data in database: %v", err.Error())
		c.Status(http.StatusUnprocessableEntity)
		return
	}

	sendSuccess(c, http.StatusNoContent, "update-opening", opening)
}
