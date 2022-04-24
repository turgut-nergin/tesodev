package response_models

type StatusError struct {
	Code int
	Err  error
}
