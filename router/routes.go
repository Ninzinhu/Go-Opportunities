package router

import (
	"github.com/Ninzinhu/Go-Opportunities/handler"
	"github.com/gin-gonic/gin"
)



func initializeRoutes(router *gin.Engine){
	v1 := router.Group("/api/v1")
	{
		// Show Opening
		v1.GET("/opening", handler.ShowOpeningHangler)
		v1.POST("/opening", handler.CreateOpeningHangler)
		v1.DELETE("/opening", handler.DeleteOpeningHangler)
		v1.PUT("/opening", handler.UpdateOpeningHangler)
		v1.GET("/openings", handler.ListOpeningHangler)
		}}
	