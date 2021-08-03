package routes

import "github.com/gin-gonic/gin"

func InitRouters(r *gin.Engine) {
	slaGroup := r.Group("/sla")
	{
		initSlaRoutes(slaGroup)
	}
}
