package responses

import "github.com/99-66/SLx-Api/models"

type SLAResponse struct {
	UpTime   models.SlaTime `json:"uptime"`
	DownTime models.SlaTime `json:"downtime"`
}
