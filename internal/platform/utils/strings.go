package utils

import (
	"strconv"
	"strings"
)

func ParseCoordinates(location string) (any, any, error) {
	coordinates := strings.Split(location, ",")
	if len(coordinates) != 2 {
		return 0, 0, nil
	}

	latitude, err := strconv.ParseFloat(coordinates[0], 64)
	if err != nil {
		return 0, 0, err
	}

	longitude, err := strconv.ParseFloat(coordinates[1], 64)
	if err != nil {
		return 0, 0, err
	}

	return latitude, longitude, nil
}
