package helpers

import (
	"net/http"
	"errors"
)

type HttpClient interface {
	Get(string) (*http.Response, error)
}

type httpClient struct{}

func (h httpClient) Get(url string) (*http.Response, error) {
	client := &http.Client{}
	response, err := client.Get(url)
	if (err != nil) {
		return response, err
	}

	if (response.StatusCode != 200) {
		return response, errors.New("Call to '" + url + "' returned with status " + response.Status)
	}

	return response, err
}

func NewHttpClient() httpClient {
	return httpClient{}
}
