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
