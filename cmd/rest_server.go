package main

import (
	"dbo-be-task/internal/adapters/presenters"
	"dbo-be-task/internal/adapters/repositories"
	"dbo-be-task/internal/adapters/rest/controllers"
	"dbo-be-task/internal/adapters/rest/middlewares"
	"dbo-be-task/internal/adapters/rest/routes"
	"dbo-be-task/internal/config"
	"dbo-be-task/internal/drivers/database"
	"dbo-be-task/internal/drivers/logging"
	"dbo-be-task/internal/drivers/rest"
	"dbo-be-task/internal/helpers"
	"dbo-be-task/internal/usecases/customer"
	"dbo-be-task/internal/usecases/order"
	"dbo-be-task/internal/usecases/user"
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"
)

func main() {
	driverConfig := config.NewDriverConfig()
	RESTServer := rest.NewGinServer(driverConfig.Server.REST)
	log := logging.NewLogger(driverConfig.Logging)
	sqlDB := database.NewSQLDatabase(driverConfig.Database, log)

	bootstrapRESTServer(&config.BootstrapConfig{
		RESTServer:   RESTServer,
		Log:          log,
		SqlDB:        sqlDB,
		DriverConfig: driverConfig,
	})

	httpServer := &http.Server{
		Addr:    fmt.Sprintf(":%d", driverConfig.Server.REST.Port),
		Handler: RESTServer,
		// ReadTimeout:  time.Duration(driverConfig.Server.REST.Timeout.Read),
		// WriteTimeout: time.Duration(driverConfig.Server.REST.Timeout.Write),
	}

	err := httpServer.ListenAndServe()

	if err != nil {
		log.WithFields(logrus.Fields{
			"error": err,
		}).Fatal("Failed to start server")
	}
}

func bootstrapRESTServer(bootstrapConfig *config.BootstrapConfig) {
	// Init helper services
	errorHelperService := helpers.NewErrorHelper(bootstrapConfig.Log)
	parserHelperService := helpers.NewParserHelper(errorHelperService)
	securityHelperService := helpers.NewSecurityHelper(bootstrapConfig.DriverConfig.Security)

	// Init repositories
	customerRepository := repositories.NewCustomerRepository(errorHelperService, bootstrapConfig.SqlDB, bootstrapConfig.Log)
	userRepository := repositories.NewUserRepository(errorHelperService, bootstrapConfig.SqlDB, bootstrapConfig.Log)
	orderRepository := repositories.NewOrderRepository(errorHelperService, bootstrapConfig.SqlDB, bootstrapConfig.Log)

	// Init presenters
	customerPresenter := presenters.NewCustomerPresenter(bootstrapConfig.Log)
	userPresenter := presenters.NewUserPresenter(bootstrapConfig.Log)
	orderPresenter := presenters.NewOrderPresenter(bootstrapConfig.Log)

	// Init usecases
	customerUsecase := customer.NewCustomerUsecase(customerRepository, customerPresenter, errorHelperService, bootstrapConfig.Log)
	userUsecase := user.NewUserUsecase(userRepository, userPresenter, securityHelperService, errorHelperService, bootstrapConfig.Log)
	orderUsecase := order.NewOrderUsecase(orderRepository, orderPresenter, errorHelperService, bootstrapConfig.Log)

	// Init controllers
	customerController := controllers.NewCustomerController(errorHelperService, parserHelperService, bootstrapConfig.Log, customerUsecase)
	userController := controllers.NewUserController(userUsecase, errorHelperService, parserHelperService, bootstrapConfig.Log)
	orderController := controllers.NewOrderController(errorHelperService, parserHelperService, bootstrapConfig.Log, orderUsecase)

	// Init middlewares
	authMiddleware := middlewares.NewAuthMiddleware(errorHelperService, bootstrapConfig.DriverConfig.Security, bootstrapConfig.Log)

	// Init app routes
	restRoute := routes.NewRESTRoute(customerController, userController, orderController, authMiddleware)

	// Setup routes
	restRoute.SetupRoutes(bootstrapConfig.RESTServer)
}
