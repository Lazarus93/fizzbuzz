package handler

import (
	"bytes"
	"encoding/json"
	"fizzbuzz/internal/fizzbuzz"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFizzBuzzHandler(t *testing.T) {
	fbService := fizzbuzz.NewService()
	h := NewHandler(fbService)
	router := gin.Default()
	router.POST("/fizzbuzz", h.FizzBuzzHandler)

	tests := []struct {
		name           string
		request        Request
		expectedStatus int
		expectedBody   string
	}{
		{
			name: "Valid request",
			request: Request{
				String1: "fizz",
				String2: "buzz",
				Int1:    3,
				Int2:    5,
				Limit:   15,
			},
			expectedStatus: http.StatusOK,
			expectedBody:   `{"result":["1","2","fizz","4","buzz","fizz","7","8","fizz","buzz","11","fizz","13","14","fizzbuzz"]}`,
		},
		{
			name: "Invalid request",
			request: Request{
				String1: "fizz",
				String2: "buzz",
				Int1:    0,
				Int2:    5,
				Limit:   15,
			},
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   `{"error":"int1 and int2 must be positive"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			body, _ := json.Marshal(tt.request)
			req, _ := http.NewRequest("POST", "/fizzbuzz", bytes.NewBuffer(body))
			resp := httptest.NewRecorder()

			router.ServeHTTP(resp, req)

			if resp.Code != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, resp.Code)
			}
			if resp.Body.String() != tt.expectedBody {
				t.Errorf("Expected body %s, got %s", tt.expectedBody, resp.Body.String())
			}
		})
	}
}
