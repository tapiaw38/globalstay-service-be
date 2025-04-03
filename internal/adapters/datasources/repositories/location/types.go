package location

type (
	LocationDocument struct {
		ID        string  `json:"id"`
		Name      string  `json:"name"`
		City      string  `json:"city"`
		State     string  `json:"state"`
		Country   string  `json:"country"`
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
		Radius    uint    `json:"radius"`
	}
)
