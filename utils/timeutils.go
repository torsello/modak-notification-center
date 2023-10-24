package timeutils

import "time"

var UnitsTime = map[string]string{
    "hour"    : "hour",
    "minute"  : "minute",
    "second"  : "second",
    "day"     : "day",
    "month"   : "month",
}

func GetValidationDate(t time.Time, timeUnit string) time.Time {
	unit, exists := UnitsTime[timeUnit]
	
    if !exists {
        return t
    }

    switch unit {
		case "hour":
			return t.Add(-1 * time.Hour)
		case "minute":
			return t.Add(-1 * time.Minute)
		case "second":
			return t.Add(-1 * time.Second)
		case "day":
			return t.AddDate(0, 0, -1)
		case "month":
			return t.AddDate(0, -1, 0)
		default:
			return t
    }
}
