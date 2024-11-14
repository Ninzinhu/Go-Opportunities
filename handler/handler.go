package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func  CreateOpeningHangler(ctx *gin.Context){
	ctx.JSON(http.StatusOK, gin.H{
		"msg" : "POST Opening",
	})

}

func  ShowOpeningHangler(ctx *gin.Context){
	ctx.JSON(http.StatusOK, gin.H{
		"msg" : "GET Opening",
	})

}

func  DeleteOpeningHangler(ctx *gin.Context){
	ctx.JSON(http.StatusOK, gin.H{
		"msg" : "GET Opening",
	})

}

func  UpdateOpeningHangler(ctx *gin.Context){
	ctx.JSON(http.StatusOK, gin.H{
		"msg" : "GET Opening",
	})

}

func  ListOpeningHangler(ctx *gin.Context){
	ctx.JSON(http.StatusOK, gin.H{
		"msg" : "GET Opening",
	})

}