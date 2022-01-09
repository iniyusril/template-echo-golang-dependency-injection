package app

import (
	"github.com/iniyusril/template/controller"
	"github.com/iniyusril/template/exception"
	cstmmiddleware "github.com/iniyusril/template/middleware"
	"github.com/labstack/echo/v4/middleware"

	"github.com/labstack/echo/v4"
)

func NewRouter(categoryController controller.CategoryController, userController controller.UserController) *echo.Echo {
	e := echo.New()
	g := e.Group("/api/categories")
	g.Use(cstmmiddleware.ServeHTTP)

	g.GET("", categoryController.FindAll)
	g.GET("/:categoryId", categoryController.FindById)
	g.POST("", categoryController.Create)
	g.PUT("/:categoryId", categoryController.Update)
	g.DELETE("/:categoryId", categoryController.Delete)

	user := e.Group("/api/v1/users")
	user.GET("", userController.FindAll)
	user.GET("/:userId", userController.FindById)
	user.POST("", userController.Create)
	user.PUT("/:userId", userController.Update)
	user.DELETE("/:userId", userController.Delete)

	e.HTTPErrorHandler = exception.ErrorHandler
	e.Use(
		middleware.Recover(),   // Recover from all panics to always have your server up
		middleware.Logger(),    // Log everything to stdout
		middleware.RequestID(), // Generate a request id on the HTTP response headers for identification
	)

	return e
}
