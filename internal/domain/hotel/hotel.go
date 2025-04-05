package domain

type (
	Hotel struct {
		ID           string
		UserID       string
		PlaceID      string
		Type         string
		Name         string
		Description  string
		PhoneNumber  string
		Email        string
		Content      string
		Address      string
		Active       bool
		Latitude     float64
		Longitude    float64
		CreatedAt    string
		UpdatedAt    string
		Rooms        []Room
		Pictures     []string
		Reviews      []Review
		AveragePrice float64
		LocationName string
	}

	Service struct {
		ID          string
		Name        string
		Description string
		Icons       []string
	}

	Review struct {
		UserName string
		Rating   float64
		Comment  string
	}

	Room struct {
		Number        string
		Type          string
		PersonCount   int
		PricePerNight float64
		Pictures      []string
		Services      []Service
	}
)

func (h Hotel) HasCompletePlaceInformation() bool {
	return h.Name != "" &&
		h.Latitude != 0 &&
		h.Longitude != 0 &&
		len(h.Pictures) > 0
}
