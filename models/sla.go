package models

import (
	"fmt"
	"strings"
	"time"
)

const (
	DAY          = 1
	SecondPerDay = 86400
	// YEAR http://en.wikipedia.org/wiki/Year#Summary
	// http://manse.ndsl.kr/contents-3.html
	// 그레고리력의 1년 평균 길이는 365.2425일이다
	YEAR  = 365.2425
	MONTH = YEAR / 12
)

// Periods 일자별로 일(day) 수를 가져오기 위한 참조 map 변수이다
var Periods = map[string]float64{
	"Daily":   DAY,
	"Weekly":  DAY * 7,
	"Monthly": MONTH,
	"Yearly":  YEAR,
}

// SLATime SLA 가동율 시간을 기간별로 저장하는 구조체이다
type SLATime struct {
	Daily   string `json:"daily"`
	Weekly  string `json:"weekly"`
	Monthly string `json:"monthly"`
	Yearly  string `json:"yearly"`
}

// NewSlaTime SLATime 구조체의 생성자 함수이다
func NewSlaTime() *SLATime {
	return &SLATime{
		Daily:   "0",
		Weekly:  "0",
		Monthly: "0",
		Yearly:  "0",
	}
}

// Calculator uptime, downtime 별로 시간을 계산하여 반환한다
// 이 함수는 timeType(uptime or downtime) 을 전달받아 해당 타입에 맞는 함수를 호출하는 래퍼 함수이다
func (s *SLATime) Calculator(percent, day float64, t string) string {
	if t == "uptime" {
		return s.upTime(percent, day)
	}

	return s.downTime(percent, day)
}

// DownTime 가동율에서 허용 가능한 downtime 을 계산하여 반환한다
// 반환 값은 1d 3h3m10s 와 같은 문자열 형식으로 반환된다
func (s *SLATime) downTime(percent, day float64) string {
	// 가동율을 분수로 변환한다
	// ex. SLA 99.9% = 99.9% / 100%
	uptimeFraction := percent / 100
	// 가동 일자를 초(Seconds) 단위로 변환한다
	SecondsPerDays := SecondPerDay * day
	// 가동 중지가 허용되는 시간을 계산한다
	downTime := SecondsPerDays - (uptimeFraction * SecondsPerDays)

	return (time.Duration(downTime) * time.Second).String()
}

// UpTime 가동율을 보장하기 위한 uptime 을 계산하여 반환한다
// 반환 값은 1d 3h13m10s 와 같은 문자열 형식으로 반환된다
func (s *SLATime) upTime(percent, day float64) string {
	// 가동율을 분수로 변환한다
	// ex. SLA 99.9% = 99.9% / 100%
	uptimeFraction := percent / 100
	// 가동 일자를 초(Seconds) 단위로 변환한다
	SecondsPerDays := SecondPerDay * day
	// 가동 해야하는 시간을 계산한다
	upTime := time.Duration(uptimeFraction*SecondsPerDays) * time.Second

	return s.fmtDuration(upTime)
}

// fmtDuration Uptime 함수 반환 값을 사용자에게 읽기 쉬운 형식으로 변환시켜준다
func (s *SLATime) fmtDuration(d time.Duration) string {
	day := time.Minute * 60 * 24
	year := 365 * day

	if d < day {
		return d.String()
	}

	var b strings.Builder

	if d >= year {
		years := d / year
		fmt.Fprintf(&b, "%dy ", years)
		d -= years * year
	}

	days := d / day
	d -= days * day
	fmt.Fprintf(&b, "%dd %s", days, d)

	return b.String()
}
