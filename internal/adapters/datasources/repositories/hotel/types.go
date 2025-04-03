package hotel

type (
	HotelDocument struct {
		ID           string           `bson:"_id,omitempty"`
		PlaceID      string           `bson:"place_id,omitempty"`
		UserID       string           `bson:"user_id,omitempty"`
		Type         string           `bson:"type,omitempty"`
		Name         string           `bson:"name,omitempty"`
		Description  string           `bson:"description,omitempty"`
		PhoneNumber  string           `bson:"phone_number,omitempty"`
		Email        string           `bson:"email,omitempty"`
		Content      string           `bson:"content,omitempty"`
		Address      string           `bson:"address,omitempty"`
		Active       bool             `bson:"active,omitempty"`
		Latitude     float64          `bson:"latitude,omitempty"`
		Longitude    float64          `bson:"longitude,omitempty"`
		CreatedAt    string           `bson:"created_at,omitempty"`
		UpdatedAt    string           `bson:"updated_at,omitempty"`
		Rooms        []RoomDocument   `bson:"rooms,omitempty"`
		Pictures     []string         `bson:"pictures,omitempty"`
		Reviews      []ReviewDocument `bson:"reviews,omitempty"`
		AveragePrice float64          `bson:"average_price,omitempty"`
		LocationName string           `bson:"location_name,omitempty"`
	}

	RoomDocument struct {
		ID            string            `bson:"_id,omitempty"`
		Number        string            `bson:"number,omitempty"`
		Type          string            `bson:"type,omitempty"`
		IsOccupied    bool              `bson:"is_occupied,omitempty"`
		PersonCount   int               `bson:"person_count,omitempty"`
		GuestName     string            `bson:"guest_name,omitempty"`
		PricePerNight float64           `bson:"price_per_night,omitempty"`
		Pictures      []string          `bson:"pictures,omitempty"`
		Services      []ServiceDocument `bson:"services,omitempty"`
	}

	ServiceDocument struct {
		ID          string   `bson:"_id,omitempty"`
		Name        string   `bson:"name,omitempty"`
		Description string   `bson:"description,omitempty"`
		Pictures    []string `bson:"pictures,omitempty"`
	}

	ReviewDocument struct {
		UserName string  `bson:"user_name,omitempty"`
		Rating   float64 `bson:"rating,omitempty"`
		Comment  string  `bson:"comment,omitempty"`
	}
)
