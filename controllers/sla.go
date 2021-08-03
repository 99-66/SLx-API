package controllers

import (
	"github.com/99-66/SLx-Api/models"
	"github.com/99-66/SLx-Api/models/consts"
	"github.com/99-66/SLx-Api/models/responses"
	"github.com/99-66/SLx-Api/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func SlaPercentage(c *gin.Context) {
	// Query Param 으로 가동율을 받는다
	strPercent := c.Param("percentage")
	// 가동율을 str -> float64 타입을 변경한다
	percentage, err := strconv.ParseFloat(strPercent, 64)
	if err != nil {
		defaultErrResponse(c, http.StatusBadRequest, consts.HTTPBadRequest, "Bad Request")
		return
	}

	// 기간별로 SLA 가동율의 가동을 보장해야 하는 시간(uptime)을 계산한다
	uptime := models.NewSlaTime()
	services.SlaCalculators(percentage, uptime, "uptime")

	// 기간별로 SLA 가동율의 가동 중지를 허용하는 시간(downtime)을 계산한다
	downTime := models.NewSlaTime()
	services.SlaCalculators(percentage, downTime, "downtime")

	response := responses.SLAResponse{
		UpTime:   *uptime,
		DownTime: *downTime,
	}

	c.JSON(http.StatusOK, response)
}
