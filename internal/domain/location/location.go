package domain

import "math"

type (
	Location struct {
		ID        string
		Name      string
		City      string
		State     string
		Country   string
		Latitude  float64
		Longitude float64
		Radius    uint
	}
)

func toRadians(deg float64) float64 {
	return deg * math.Pi / 180
}

func (l Location) haversineDistanceTo(lat, lon float64) float64 {
	const R = 6371
	lat1, lon1, lat2, lon2 := toRadians(l.Latitude), toRadians(l.Longitude), toRadians(lat), toRadians(lon)

	dLat := lat2 - lat1
	dLon := lon2 - lon1

	a := math.Sin(dLat/2)*math.Sin(dLat/2) +
		math.Cos(lat1)*math.Cos(lat2)*
			math.Sin(dLon/2)*math.Sin(dLon/2)

	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	return R * c
}

func FindNearestLocation(targetLat, targetLon float64, locations []Location) Location {
	var nearest Location
	minDistance := math.MaxFloat64

	for _, loc := range locations {
		distance := loc.haversineDistanceTo(targetLat, targetLon)
		if distance < minDistance {
			minDistance = distance
			nearest = loc
		}
	}

	return nearest
}
