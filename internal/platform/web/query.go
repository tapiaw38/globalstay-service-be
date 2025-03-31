package web

import (
	"strconv"
	"strings"
	"time"
)

func ParseStringArrayQueryValue(value string, delimiter string) []string {
	parsedValue := strings.Split(value, delimiter)
	if len(parsedValue) == 1 && parsedValue[0] == "" {
		return []string{}
	}

	return parsedValue
}

func ParseInt64QueryValue(value string) int64 {
	int64Value, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return 0
	}

	return int64Value
}

func ParseTimeQueryValue(value string, layout string) time.Time {
	timeValue, err := time.Parse(layout, value)
	if err != nil {
		return time.Time{}
	}

	return timeValue
}

func ParseBoolQueryValue(value string) bool {
	boolValue, err := strconv.ParseBool(value)
	if err != nil {
		return false
	}

	return boolValue
}

func ParseBoolPointerQueryValue(value string) *bool {
	boolValue := ParseBoolQueryValue(value)
	if boolValue {
		return &boolValue
	}

	return nil
}

func ParseFloat64QueryValue(value string) float64 {
	float64Value, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return 0
	}

	return float64Value
}
