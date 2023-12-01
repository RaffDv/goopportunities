package router

import (
	"github.com/RaffDv/goopportunities/internal/handler"
	"github.com/gin-gonic/gin"
)

func initializeRoutes(r *gin.Engine) {
	v1 := r.Group("/api/v1")

	opp := v1.Group("/opening")
	{
		opp.GET("", handler.FindOpeningHandler)
		opp.GET("/all", handler.ListOpeningsHandler)
		opp.POST("", handler.CreateOpeningHandler)
		opp.DELETE("", handler.DeleteOpeningHandler)
		opp.PUT("", handler.UpdateOpeningHandler)
	}

}
