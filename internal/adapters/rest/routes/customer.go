package routes

import (
	"github.com/gin-gonic/gin"
)

func (rt *RESTRoute) AttachCustomerRoutesV1(router *gin.RouterGroup) {
	router.GET("/customers", rt.CustomerController.GetCustomers)
	router.GET("/customers/:id", rt.CustomerController.GetCustomerByID)
	router.POST("/customers", rt.CustomerController.UpsertCustomer)
	router.PUT("/customers/:id", rt.CustomerController.UpdateCustomerByID)
	router.DELETE("/customers/:id", rt.CustomerController.DeleteCustomerByID)
}
