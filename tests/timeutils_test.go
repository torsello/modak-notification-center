package tests

import (
	myUtils "modak-notification-center/utils"
	"testing"
	"time"
)

func TestGetValidationDate(t *testing.T) {
	referenceTime := time.Date(2023, time.October, 24, 12, 0, 0, 0, time.UTC)

	tests := []struct {
		unit       string
		expected   time.Time
	}{
		{"hour", referenceTime.Add(-1 * time.Hour)},
		{"minute", referenceTime.Add(-1 * time.Minute)},
		{"second", referenceTime.Add(-1 * time.Second)},
		{"day", referenceTime.AddDate(0, 0, -1)},
		{"month", referenceTime.AddDate(0, -1, 0)},
	}

	for _, test := range tests {
		t.Run(test.unit, func(t *testing.T) {
			result := myUtils.GetValidationDate(referenceTime, test.unit)
			if !result.Equal(test.expected) {
				t.Errorf("For the time unit %s, %v was expected but %v was obtained", test.unit, test.expected, result)
			}
		})
	}
}
