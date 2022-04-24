package response_models

type Product struct {
	ProductId string `json:"productid"`
	ImageUrl  string `json:"imageurl"`
	Name      string `json:"name"`
}
