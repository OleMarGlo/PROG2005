package handler

type Coordinates struct {
	Latitude   float64 `json:"latitude"`
	Longtitude float64 `json:"longtitude"`
}

type Location struct {
	Name        string      `json:"name"`
	Postcode    string      `json:"code"`
	Country     string      `json:"country,omitempty"`
	Geolocation Coordinates `json:"geolocation"`
}
