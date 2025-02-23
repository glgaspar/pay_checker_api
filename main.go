package main

import (
	"github.com/gin-gonic/gin"
	"github.com/glgaspar/pay_checker_api/controller"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Printf("Error loading .env file: %s", err)
	}

	router := gin.Default()

	setUpRoutes(router)

	router.Run("0.0.0.0:8080")
}

func setUpRoutes(router *gin.Engine) {
	// healthcheck
	router.GET("/", func(c *gin.Context) { c.Done() })

	// create new bill
	router.POST("/", controller.CreateBill)

	// update bill data
	router.POST("/udpate", controller.UpdateBill)
	router.POST("/pay/:billId", controller.PayBill)

	// list bills
	router.GET("/list", controller.GetList)
}
