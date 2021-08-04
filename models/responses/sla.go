package responses

import "github.com/99-66/SLx-Api/models"

type SLAResponse struct {
	//UpTime   models.SLATime `json:"uptime"`
	DownTime models.SLATime `json:"downtime"`
}
