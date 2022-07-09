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
			t.Fatalf("FAILED: expected status code %d; got %d", http.StatusOK, response.StatusCode)
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
		{name: "request with incorrect endpoint", endpoint: "/books/1", httpStatusCode: 404},
		{name: "request with spelling error", endpoint: "/boojs", httpStatusCode: 404},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			testHandler := func() http.Handler {
				r := http.NewServeMux()
				r.HandleFunc(tc.endpoint, GetBook)
				return r
			}
			srv := httptest.NewServer(testHandler())
			defer srv.Close()

			response, err := http.Get(fmt.Sprintf("%s/books", srv.URL))
			if err != nil {
				t.Fatalf("could not send GET request: %v", err)
			}

			if response.StatusCode != tc.httpStatusCode {
				t.Fatalf("FAILED: expected status code %d; got %d", tc.httpStatusCode, response.StatusCode)
			}
		})

	}

}

// func TestGetBookById(t *testing.T) {
// 	testcases := []struct {
// 		name           string
// 		endpoint       string
// 		httpStatusCode int
// 		contentLength  int64
// 	}{
// 		{name: "book exists in database", endpoint: fmt.Sprintf("/books/%d", 7), httpStatusCode: 200},
// 		{name: "book does not exist in database", endpoint: "/books/44", httpStatusCode: 404, contentLength: 74},
// 		{name: "request with spelling error", endpoint: "/book/7", httpStatusCode: 404, contentLength: 74},
// 	}

// 	for _, tc := range testcases {
// 		t.Run(tc.name, func(t *testing.T) {
// 			testHandler := func() http.Handler {
// 				r := http.NewServeMux()
// 				r.HandleFunc(tc.endpoint, GetBookById)
// 				return r
// 			}

// 			srv := httptest.NewServer(testHandler())
// 			defer srv.Close()

// 			response, err := http.Get(fmt.Sprintf("%s/books", srv.URL))
// 			if err != nil {
// 				t.Fatalf("could not send GET request: %v", err)
// 			}
// 			// fmt.Println(response)
// 			if response.StatusCode != tc.httpStatusCode {
// 				t.Fatalf("FAILED: expected status code %d; got %d", tc.httpStatusCode, response.StatusCode)
// 			}

// 			if response.ContentLength != tc.contentLength {
// 				t.Fatalf("FAILED: expected response with Content-Length %d; got %d", tc.contentLength, response.ContentLength)
// 			}
// 		})
// 	}
// }
