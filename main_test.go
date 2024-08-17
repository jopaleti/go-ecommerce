package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/jopaleti/go-ecommerce/helpers"
	"github.com/stretchr/testify/assert"
)

func TestHomeHandler(t *testing.T) {
	mockResponse := `{"message":"Welcome to E-Commerce API"}`

	r := helpers.SetUpRouter()
	r.GET("/", func (c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to E-Commerce API",
		})
	})

	// Create a new HTTP GET request
	req, _ := http.NewRequest("GET", "/", nil)

	// Create a new HTTP response recorder
	w := httptest.NewRecorder()

	// Serve the HTTP request using the router
	r.ServeHTTP(w, req)

	// Read the response data from the recorder
	responseData, _ := io.ReadAll(w.Body)

	assert.Equal(t, mockResponse, string(responseData))

	assert.Equal(t, http.StatusOK, w.Code)
}