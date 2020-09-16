package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/samherrmann/serveit/handlers"
)

func TestFileHandler(t *testing.T) {
	type Want struct {
		status int
		body   string
	}

	type Test struct {
		notFoundFile  string
		requestedFile string
		want          *Want
	}

	// Define test cases
	tests := []Test{
		{
			notFoundFile:  "",
			requestedFile: "/testdata/request-1.txt",
			want: &Want{
				status: http.StatusOK,
				body:   "This file is used for testing.",
			},
		}, {
			notFoundFile:  "",
			requestedFile: "/some-file-that-does-not-exist.txt",
			want: &Want{
				status: http.StatusNotFound,
				body:   "404 page not found\n",
			},
		}, {
			notFoundFile:  "./testdata/custom-404.txt",
			requestedFile: "/some-file-that-does-not-exist.txt",
			want: &Want{
				status: http.StatusNotFound,
				body:   "Custom 404 page.",
			},
		},
	}

	// Loop over all test cases
	for _, tc := range tests {
		// Create a request to pass to the handler.
		req, err := http.NewRequest("GET", tc.requestedFile, nil)
		if err != nil {
			t.Fatal(err)
		}

		// Setup handler with response recorder
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(handlers.FileHandler(tc.notFoundFile))
		handler.ServeHTTP(rr, req)

		// Check the status code is what we expect.
		if status := rr.Code; status != tc.want.status {
			t.Errorf(
				"Handler returned wrong status code: want %v, but got %v.",
				tc.want.status,
				status,
			)
		}

		// Check the response body is what we expect.
		if body := rr.Body.String(); body != tc.want.body {
			t.Errorf(
				"Handler returned unexpected body: want \"%v\", but got \"%v\".",
				tc.want.body,
				body,
			)
		}
	}
}
