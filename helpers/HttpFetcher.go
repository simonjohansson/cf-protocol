package helpers

import (
	"net/http"
	"errors"
	"io/ioutil"
	"bytes"
)

type HttpClient interface {
	Get(string) (*http.Response, error)
}

// Impl
//
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

// Mock
//
type mockHttpClient struct {
	Request map[string]*http.Response
}

func (m mockHttpClient) SetupTestData(url string, body string, status int) {
	response := &http.Response{
		Body:       ioutil.NopCloser(bytes.NewBuffer([]byte(body))),
		StatusCode: status,
	}
	m.Request[url] = response
}

func NewMockHttpFetcher() mockHttpClient {
	return mockHttpClient{
		Request: map[string]*http.Response{},
	}
}

func (m mockHttpClient) Get(url string) (*http.Response, error) {
	response := m.Request[url]
	if ( response.StatusCode != 200 ) {
		return nil, errors.New("Shit")

	}
	return response, nil
}
