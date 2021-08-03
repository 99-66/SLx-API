package services

import (
	"github.com/99-66/SLx-Api/models"
	"reflect"
)

// SlaCalculators SLA Percentage 를 계산하여 Uptime/Downtime 을 반환한다
// Uptime: SLA Percentage 를 만족하는 서비스가 동작해야 하는 시간
// Downtime: SLA Percentage 를 만족하는 서비스가 중지가 허용되는 시간
func SlaCalculators(percent float64, sla *models.SLATime, t string) {
	// Structure iterate 를 위한 refection
	v := reflect.ValueOf(sla).Elem()

	// SlaTime 구조체 변수를 순회하며 SLA 시간을 계산한다
	for i := 0; i < v.NumField(); i++ {
		// Structure 의 필드 이름을 동적으로 가져온다
		fieldName := v.Type().Field(i).Name
		// 동적 필드 값별로 slaTime 을 계산한다
		slaTime := sla.Calculator(percent, models.Periods[fieldName], t)
		// 동적 필드 값별로 계산한 SLA 값을 업데이트한다
		v.Field(i).Set(reflect.ValueOf(slaTime))
	}
}
