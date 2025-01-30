package router

import "github.com/gin-gonic/gin"

func NewRouter() {
	router := gin.Default()

	newRoutes(router)

	router.Run(":8080")
}
