package response_models

type Address struct {
	AddressLine string `json:"addressline"`
	City        string `json:"city"`
	Country     string `json:"country"`
	CityCode    int    `json:"citycode"`
}
