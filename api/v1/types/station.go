package types

type StationResponse struct {
	Ok      bool   `json:"ok"`
	License string `json:"license"`
	Data    string `json:"data"`
	Status  string `json:"status"`
	Station struct {
		ID           string `json:"id"`
		Name         string `json:"name"`
		Brand        string `json:"brand"`
		Street       string `json:"street"`
		HouseNumber  string `json:"houseNumber"`
		PostCode     int    `json:"postCode"`
		Place        string `json:"place"`
		OpeningTimes []struct {
			Text  string `json:"text"`
			Start string `json:"start"`
			End   string `json:"end"`
		} `json:"openingTimes"`
		Overrides []interface{} `json:"overrides"`
		WholeDay  bool          `json:"wholeDay"`
		IsOpen    bool          `json:"isOpen"`
		E5        float64       `json:"e5"`
		E10       float64       `json:"e10"`
		Diesel    float64       `json:"diesel"`
		Lat       float64       `json:"lat"`
		Lng       float64       `json:"lng"`
		State     interface{}   `json:"state"`
	} `json:"station"`
}
