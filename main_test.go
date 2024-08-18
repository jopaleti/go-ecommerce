package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/jopaleti/go-ecommerce/controllers"
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

type userModel struct {
	Email	string	`json:"email"      validate:"email,required"`
	Password string	`json:"password"   validate:"required,min=6"`
	First_Name string `json:"first_name"  validate:"required"`
	Last_Name string `json:"last_name"   validate:"required"`
	Phone string `json:"phone" validate:"required"`
}

func TestLogin (t *testing.T) {

	r := helpers.SetUpRouter()
	r.POST("/users/login", controllers.Login())

	user := userModel{
		Email: "segunopaleti@gmail.com",
		Password: "tobi1234",
	}
	
	jsonValue, _ := json.Marshal(user)
	req, _ := http.NewRequest("POST", "/users/login", bytes.NewBuffer(jsonValue))

	w := httptest.NewRecorder()
	
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusFound, w.Code)
}

func TestSignUp (t *testing.T) {

	r := helpers.SetUpRouter()
	r.POST("/users/signup", controllers.SignUp())

	user := userModel {
		Email: "oopalet7pi@gmail.com",
		Password: "tb787945",
		First_Name: "Oluwasegun",
		Last_Name: "Opaleti",
		Phone: "08345678",
	}

	jsonValue, _ := json.Marshal(user)
	req, _ := http.NewRequest("POST", "/users/signup", bytes.NewBuffer(jsonValue))

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)
}

type productModel struct {
	Product_Name string `json:"product_name"`
	Price        uint64 `json:"price"`
	Rating    uint8 `json:"rating"`
	Image     string `json:"image"`
}
func TestProductViewerAdmin(t *testing.T) {
	r := helpers.SetUpRouter()
	r.POST("/admin/addproduct", controllers.ProductViewerAdmin())

	product := productModel{
		Product_Name: "Samsung Galaxy S21",
		Price: 200000,
		Rating: 4,
		Image: "https://www.google.com",
	}

	jsonValue, _ := json.Marshal(product)
	req, _ := http.NewRequest("POST", "/admin/addproduct", bytes.NewBuffer(jsonValue))

	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)
	assert.Equal(t, http.StatusCreated, response.Code)
}