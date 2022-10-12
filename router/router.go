package router

import (
	"Assignment2/controllers"
	"Assignment2/database"
	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	router := gin.Default()

	db := database.GetDB()
	//db.AutoMigrate(models.Order{}, models.Item{})
	orderController := controllers.NewOrderController(db)

	orderGroup := router.Group("/orders")
	{
		orderGroup.POST("/", orderController.CreateOrder)
		orderGroup.GET("/", orderController.GetOrders)
		orderGroup.PUT("/:orderId", orderController.UpdateOrder)
		orderGroup.DELETE("/:orderId", orderController.DeleteOrder)
	}

	return router
}
