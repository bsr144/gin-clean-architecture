package routes

import (
	"dbo-be-task/internal/adapters/rest/controllers"
	"dbo-be-task/internal/adapters/rest/middlewares"

	"github.com/gin-gonic/gin"
)

type RESTRoute struct {
	AuthMiddleware     *middlewares.AuthMiddleware
	CustomerController *controllers.CustomerController
	UserController     *controllers.UserController
	OrderController    *controllers.OrderController
}

func NewRESTRoute(customerController *controllers.CustomerController, userController *controllers.UserController, orderController *controllers.OrderController, authMiddleware *middlewares.AuthMiddleware) *RESTRoute {
	return &RESTRoute{
		CustomerController: customerController,
		UserController:     userController,
		OrderController:    orderController,
		AuthMiddleware:     authMiddleware,
	}
}

func (rt *RESTRoute) SetupRoutes(server *gin.Engine) {
	api := server.Group("/api")
	v1 := api.Group("/v1")

	rt.SetupPublicRoutesV1(v1)
	rt.SetupPrivateRoutesV1(v1)
}

func (rt *RESTRoute) SetupPublicRoutesV1(router *gin.RouterGroup) {
	rt.AttachAuthRoutesV1(router)
}

func (rt *RESTRoute) SetupPrivateRoutesV1(router *gin.RouterGroup) {
	router.Use(rt.AuthMiddleware.VerifyAuth)
	rt.AttachCustomerRoutesV1(router)
	rt.AttachOrderRoutesV1(router)
}
