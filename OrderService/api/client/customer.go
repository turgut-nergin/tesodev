package client

import (
	"errors"
	"io"
	"net/http"
	"net/url"
	"path"

	"github.com/turgut-nergin/tesodev/api/handlers/response_models"
)

var baseURL = url.URL{
	Scheme: "http",
	Host:   "customer-service:8086",
	Path:   "/validate/customer/",
}

func Get(customerId string) (bool, *response_models.StatusError) {
	var statusError *response_models.StatusError

	endpt := baseURL.ResolveReference(
		&url.URL{Path: "customerId"})

	endpt.Path = path.Join(endpt.Path, customerId)
	res, err := http.Get(endpt.String())

	if err != nil {
		panic(err)
	}

	response, err := io.ReadAll(res.Body)
	if err != nil {
		statusError.Code = http.StatusBadRequest
		statusError.Err = errors.New(string(response))
	}

	if res.StatusCode != http.StatusOK {
		statusError.Code = res.StatusCode
		statusError.Err = errors.New(string(response))
		return false, statusError
	}

	return string(response) == "true", nil

}
