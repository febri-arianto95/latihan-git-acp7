package routes

import (
	"project/constants"
	"project/controllers"
	mid "project/middleware"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func New() *echo.Echo {
	e := echo.New()
	mid.LogMiddleware(e)
	e.POST("/users", controllers.CreateUsersController)
	e.POST("/login", controllers.LoginUsersController)
	e.GET("/product", controllers.GetProductController)
	// e.GET("/users", controllers.GetUsersController)

	jwt := e.Group("jwt")
	jwt.Use(middleware.JWT([]byte(constants.KEY_JWT)))
	jwt.GET("", controllers.GetUsersController)
	jwt.POST("/product", controllers.CreateProductController)
	jwt.GET("/product", controllers.GetProductController)
	jwt.POST("/cart", controllers.CreateCartController)
	jwt.GET("/cart", controllers.GetCartController)
	jwt.DELETE("/cart/:id", controllers.DeleteCartController)
	jwt.POST("/checkout", controllers.CreateTransactionController)

	e.GET("/users/companies", controllers.GetUsersCompanyController)
	e.POST("/users/companies", controllers.CreateUserCompaniesController)
	return e
}
