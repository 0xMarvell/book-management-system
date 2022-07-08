package controllers

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHomePage(t *testing.T) {
	request, err := http.NewRequest("GET", "localhost:8080/", nil)
	if err != nil {
		t.Fatalf("could not create request: %v", err)
	}

	recorder := httptest.NewRecorder()
	HomePage(recorder, request)

	res := recorder.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Fatalf("FAILED: expected Status OK (200); got %v", res.StatusCode)
	}

}

func TestGetBook(t *testing.T) {
	// request, err := http.NewRequest("GET", "localhost:8080/books", nil)
	// if err != nil {
	// 	t.Fatalf("could not create request: %v", err)
	// }

	// recorder := httptest.NewRecorder()
	// GetBook(recorder, request)

	// res := recorder.Result()
	// defer res.Body.Close()

	// if res.StatusCode != http.StatusOK {
	// 	t.Fatalf("FAILED: expected Status OK (200); got %v", res.StatusCode)
	// }
	testcases := []struct {
		name           string
		endpoint       string
		httpStatusCode int
	}{
		{name: "REQUEST WITH CORRECT ENDPOINT", endpoint: "/books", httpStatusCode: 200},
		{name: "REQUEST WITH INCORRECT ENDPOINT", endpoint: "/book/1", httpStatusCode: 404},
		{name: "REQUEST WITH SPELLING ERROR", endpoint: "/boojs", httpStatusCode: 404},
	}

	for _, tt := range testcases {
		t.Run(tt.name, func(t *testing.T) {
			testHandler := func() http.Handler {
				r := http.NewServeMux()
				r.HandleFunc(tt.endpoint, GetBook)
				return r
			}
			srv := httptest.NewServer(testHandler())
			defer srv.Close()

			response, err := http.Get(fmt.Sprintf("%s/books", srv.URL))
			if err != nil {
				t.Fatalf("could not send GET request: %v", err)
			}

			if response.StatusCode != tt.httpStatusCode {
				t.Fatalf("FAILED: expected status code %d ; got %d", tt.httpStatusCode, response.StatusCode)
			}
		})
	}

}
