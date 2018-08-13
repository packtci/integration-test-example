package main

import (
	"errors"
	"net/http"
	"os"
	"testing"

	"gopkg.in/h2non/baloo.v3"
)

var test = baloo.New("https://api.travis-ci.com")

func assertTravisUserEndpoint(res *http.Response, req *http.Request) error {
	if res.StatusCode != http.StatusOK {
		return errors.New("This endpoint should return a 200 response code")
	}
	if res.Body == nil {
		return errors.New("The body should not be empty")
	}
	return nil
}

func TestBalooClient(t *testing.T) {
	test.Get("/user").
		SetHeader("Authorization", "token "+os.Getenv("TRAVIS_PERSONAL_TOKEN")).
		SetHeader("Travis-API-Version", "3").
		Expect(t).
		Status(200).
		Type("json").
		AssertFunc(assertTravisUserEndpoint).
		Done()
}
