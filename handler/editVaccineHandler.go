package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func EditVaccineHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Edit Vaccine!",
	})
}
