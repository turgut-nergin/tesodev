package response_models

type Error struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
}
