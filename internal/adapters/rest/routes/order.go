package routes

import (
	"github.com/gin-gonic/gin"
)

func (rt *RESTRoute) AttachOrderRoutesV1(router *gin.RouterGroup) {
	router.GET("/orders", rt.OrderController.GetOrders)
	router.GET("/orders/:id", rt.OrderController.GetOrderByID)
	router.POST("/orders", rt.OrderController.CreateOrder)
	router.PUT("/orders/:id", rt.OrderController.UpdateOrderByID)
	router.DELETE("/orders/:id", rt.OrderController.DeleteOrderByID)
}
