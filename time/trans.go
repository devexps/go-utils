package util

import (
	"time"

	"github.com/devexps/go-utils/trans"
)

var DefaultTimeLocation *time.Location

func RefreshDefaultTimeLocation(name string) {
	DefaultTimeLocation, _ = time.LoadLocation(name)
}

// UnixMilliToStringPtr millisecond timestamp -> string
func UnixMilliToStringPtr(tm *int64) *string {
	if tm == nil {
		return nil
	}
	str := time.UnixMilli(*tm).Format(TimeLayout)
	return &str
}

// StringToUnixMilliInt64Ptr string -> millisecond timestamp
func StringToUnixMilliInt64Ptr(tm *string) *int64 {
	theTime := StringTimeToTime(tm)
	if theTime == nil {
		return nil
	}
	unixTime := theTime.UnixMilli()
	return &unixTime
}

// StringTimeToTime time string -> time
func StringTimeToTime(str *string) *time.Time {
	if str == nil {
		return nil
	}
	if len(*str) == 0 {
		return nil
	}
	if DefaultTimeLocation == nil {
		RefreshDefaultTimeLocation(DefaultTimeLocationName)
	}
	var err error
	var theTime time.Time

	theTime, err = time.ParseInLocation(TimeLayout, *str, DefaultTimeLocation)
	if err == nil {
		return &theTime
	}
	theTime, err = time.ParseInLocation(DateLayout, *str, DefaultTimeLocation)
	if err == nil {
		return &theTime
	}
	theTime, err = time.ParseInLocation(ClockLayout, *str, DefaultTimeLocation)
	if err == nil {
		return &theTime
	}
	return nil
}

// TimeToTimeString time -> time string
func TimeToTimeString(tm *time.Time) *string {
	if tm == nil {
		return nil
	}
	return trans.String(tm.Format(TimeLayout))
}

// StringDateToTime string -> time
func StringDateToTime(str *string) *time.Time {
	if str == nil {
		return nil
	}
	if len(*str) == 0 {
		return nil
	}
	if DefaultTimeLocation == nil {
		RefreshDefaultTimeLocation(DefaultTimeLocationName)
	}
	var err error
	var theTime time.Time

	theTime, err = time.ParseInLocation(TimeLayout, *str, DefaultTimeLocation)
	if err == nil {
		return &theTime
	}
	theTime, err = time.ParseInLocation(DateLayout, *str, DefaultTimeLocation)
	if err == nil {
		return &theTime
	}
	theTime, err = time.ParseInLocation(ClockLayout, *str, DefaultTimeLocation)
	if err == nil {
		return &theTime
	}
	return nil
}

// TimeToDateString time -> date string
func TimeToDateString(tm *time.Time) *string {
	if tm == nil {
		return nil
	}
	return trans.String(tm.Format(DateLayout))
}
