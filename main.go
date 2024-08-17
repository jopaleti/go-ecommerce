package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jopaleti/go-ecommerce/controllers"
	"github.com/jopaleti/go-ecommerce/database"
	"github.com/jopaleti/go-ecommerce/middleware"
	"github.com/jopaleti/go-ecommerce/routes"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8001"
	}

	app := controllers.NewApplication(database.ProductData(database.Client, "Products"), database.UserData(database.Client, "Users"))

	router := gin.New()
	router.Use(gin.Logger())

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to E-Commerce API",
		})
	})
	routes.UserRoutes(router)
	router.Use(middleware.Authentication())

	router.GET("/addtocart", app.AddToCart())
	router.GET("/removeitem", app.RemoveItem())
	router.GET("/cartcheckout", app.BuyFromCart())
	router.GET("/instantbuy", app.InstantBuy())

	log.Fatal(router.Run(":" + port))
}