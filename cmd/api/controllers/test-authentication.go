package controllers

import (
	"authentification/cmd/api/helpers"
	"authentification/data/models"
	"bytes"
	"io"
	"net/http"
)

type RoundTripFunc func(req *http.Request) *http.Response

func (f RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}

func NewTestClient(fn RoundTripFunc) *http.Client {
	return &http.Client{
		Transport: RoundTripFunc(fn),
	}
}
func NewTestAuthentication() Authentication {
	jsonToReturn := `
{
	error: false,
	message: "user authenticated",
}`

	client := NewTestClient(func(req *http.Request) *http.Response {
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(bytes.NewBufferString(jsonToReturn)),
			Header:     make(http.Header),
		}
	})
	return &authenticationController{
		userModels: models.NewTestRepository(nil),
		json:       &helpers.JsonResponse{},
		Client:     client,
	}
}
