package request_models

type Address struct {
	AddressLine string `json:"addressLine"`
	City        string `json:"city"`
	Country     string `json:"country"`
	CityCode    int    `json:"cityCode"`
}
