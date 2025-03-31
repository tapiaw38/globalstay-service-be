package service

type (
	ServiceDocument struct {
		ID          string   `bson:"_id,omitempty"`
		Name        string   `bson:"name,omitempty"`
		Description string   `bson:"description,omitempty"`
		Pictures    []string `bson:"pictures,omitempty"`
	}
)
