package location

type (
	LocationDocument struct {
		ID        string  `bson:"_id,omitempty"`
		Name      string  `bson:"name,omitempty"`
		City      string  `bson:"city,omitempty"`
		State     string  `bson:"state,omitempty"`
		Country   string  `bson:"country,omitempty"`
		Latitude  float64 `bson:"latitude,omitempty"`
		Longitude float64 `bson:"longitude,omitempty"`
		Radius    uint    `bson:"radius,omitempty"`
	}
)
