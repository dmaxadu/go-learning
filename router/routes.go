package router

import (
	"github.com/dmaxadu/go-learning/handler"
	"github.com/gin-gonic/gin"
)

func newRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		v1.GET("/vaccine", handler.GetVaccineHandler)

		v1.POST("/vaccine/create", handler.AddVaccineHandler)

		v1.PUT("/vaccine/update", handler.EditVaccineHandler)

		v1.DELETE("/vaccine/delete", handler.DeleteVaccineHandler)

		v1.GET("/vaccines", handler.GetAllVaccinesHandler)
	}
}
