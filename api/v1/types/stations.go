package types

type StationsResponse struct {
	Ok       bool   `json:"ok"`
	License  string `json:"license"`
	Data     string `json:"data"`
	Status   string `json:"status"`
	Stations []struct {
		ID          string  `json:"id"`
		Name        string  `json:"name"`
		Brand       string  `json:"brand"`
		Street      string  `json:"street"`
		Place       string  `json:"place"`
		Lat         float64 `json:"lat"`
		Lng         float64 `json:"lng"`
		Dist        float64 `json:"dist"`
		Price       float64 `json:"price"`
		IsOpen      bool    `json:"isOpen"`
		HouseNumber string  `json:"houseNumber"`
		PostCode    int     `json:"postCode"`
	} `json:"stations"`
}
