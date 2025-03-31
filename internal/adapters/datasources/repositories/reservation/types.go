package reservation

type (
	ReservationDocument struct {
		ID       string             `json:"id"`
		Business BusinessDocument   `json:"business"`
		Services []ServiceDocument  `json:"services"`
		Schedule []ScheduleDocument `json:"schedule"`
	}

	ScheduleDocument struct {
		StartDate      string `json:"start_date"`
		EndDate        string `json:"end_date"`
		ReservationDay string `json:"reservation_day"`
	}

	ServiceDocument struct {
		ID          string   `json:"id"`
		BusinessID  string   `json:"business_id"`
		Name        string   `json:"name"`
		Description string   `json:"description"`
		Pictures    []string `json:"pictures"`
	}

	BusinessDocument struct {
		ID          string  `json:"id"`
		UserID      string  `json:"user_id"`
		Type        string  `json:"type"`
		Name        string  `json:"name"`
		Description string  `json:"description"`
		PhoneNumber string  `json:"phone_number"`
		Email       string  `json:"email"`
		Content     string  `json:"content"`
		Address     string  `json:"address"`
		Active      bool    `json:"active"`
		Latitude    float64 `json:"latitude"`
		Longitude   float64 `json:"longitude"`
		CreatedAt   string  `json:"created_at"`
		UpdatedAt   string  `json:"updated_at"`
	}
)
