package domain

type (
	Business struct {
		ID           string
		UserID       string
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
		Services     []Service
		Pictures     []string
		Reviews      []Review
		AveragePrice float64
	}

	Service struct {
		ID          string
		Name        string
		Description string
		Pictures    []string
	}

	Review struct {
		UserName string
		Rating   float64
		Comment  string
	}
)
