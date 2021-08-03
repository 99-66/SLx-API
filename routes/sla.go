package routes

import (
	"github.com/99-66/SLx-Api/controllers"
	"github.com/gin-gonic/gin"
)

func initSlaRoutes(rg *gin.RouterGroup) {
	rg.GET("/:percentage", controllers.SlaPercentage)
}
