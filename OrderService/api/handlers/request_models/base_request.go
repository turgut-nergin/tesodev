package request_models

type BaseRequest interface {
	Validate() error
}
