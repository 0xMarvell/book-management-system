package controllers

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHomePage(t *testing.T) {
	t.Run("load homepage with correct status code", func(t *testing.T) {
		testHandler := func() http.Handler {
			r := http.NewServeMux()
			r.HandleFunc("/", HomePage)
			return r
		}
		srv := httptest.NewServer(testHandler())
		defer srv.Close()

		response, err := http.Get(srv.URL)
		if err != nil {
			t.Fatalf("could not send GET request: %v", err)
		}

		if response.StatusCode != http.StatusOK {
			t.Fatalf("FAILED: expected status code %d ; got %d", http.StatusOK, response.StatusCode)
		}
	})

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
		{name: "request with correct endpoint", endpoint: "/books", httpStatusCode: 200},
		{name: "request with incorrect endpoint", endpoint: "/book/1", httpStatusCode: 404},
		{name: "request with spelling error", endpoint: "/boojs", httpStatusCode: 404},
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
