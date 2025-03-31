package business

type (
	BusinessDocument struct {
		ID          string  `bson:"_id,omitempty"`
		UserID      string  `bson:"user_id,omitempty"`
		Type        string  `bson:"type,omitempty"`
		Name        string  `bson:"name,omitempty"`
		Description string  `bson:"description,omitempty"`
		PhoneNumber string  `bson:"phone_number,omitempty"`
		Email       string  `bson:"email,omitempty"`
		Content     string  `bson:"content,omitempty"`
		Address     string  `bson:"address,omitempty"`
		Active      bool    `bson:"active,omitempty"`
		Latitude    float64 `bson:"latitude,omitempty"`
		Longitude   float64 `bson:"longitude,omitempty"`
		CreatedAt   string  `bson:"created_at,omitempty"`
		UpdatedAt   string  `bson:"updated_at,omitempty"`
	}
)
