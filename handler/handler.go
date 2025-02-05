package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddVaccineHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Added Vaccine!",
	})
}

func EditVaccineHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Edit Vaccine!",
	})
}

func DeleteVaccineHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Delete Vaccine!",
	})
}

func GetAllVaccinesHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Getting All Vaccine!",
	})
}

func GetVaccineHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Getting Vaccine!",
	})
}
