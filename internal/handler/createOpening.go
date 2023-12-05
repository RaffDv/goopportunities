package handler

import (
	"net/http"

	"github.com/RaffDv/goopportunities/internal/schema"
	"github.com/gin-gonic/gin"
)

func CreateOpeningHandler(c *gin.Context) {

	body := CreateOpeningRequest{}

	c.BindJSON(&body)

	if err := body.Validate(); err != nil {
		logger.Errorf(err.Error())
		sendError(c, http.StatusBadRequest, err.Error())
		return
	}

	opening := schema.Opening{
		Role:     body.Role,
		Company:  body.Company,
		Location: body.Location,
		Remote:   *body.Remote,
		Link:     body.Link,
		Salary:   body.Salary,
	}

	if err := db.Create(&opening).Error; err != nil {
		logger.Errorf("Error creating data in database: %v", err.Error())
		c.Status(http.StatusUnprocessableEntity)
		return
	}

	sendSuccess(c, http.StatusCreated, "createOpening", opening)
}
