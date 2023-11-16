package util

import "time"

const (
	DateLayout  = "2006-01-02"
	ClockLayout = "15:04:05"
	TimeLayout  = DateLayout + " " + ClockLayout

	DefaultTimeLocationName = "Asia/Ho_Chi_Minh"
)

var ReferenceTimeValue time.Time = time.Date(2006, 1, 2, 15, 4, 5, 999999999, time.FixedZone("GMT", 7*60*60))
