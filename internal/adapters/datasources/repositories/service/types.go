package service

type (
	ServiceDocument struct {
		ID          string   `bson:"_id,omitempty"`
		BusinessID  string   `bson:"business_id,omitempty"`
		Name        string   `bson:"name,omitempty"`
		Description string   `bson:"description,omitempty"`
		Pictures    []string `bson:"pictures,omitempty"`
	}
)
