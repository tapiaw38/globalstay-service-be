package hotel

type (
	ScanPlacesInput struct {
		Location string `json:"location"`
		Radius   uint   `json:"radius"`
		Types    string `json:"types"`
	}
)
