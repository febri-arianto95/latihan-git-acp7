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
	// batasi akses yang sudah punya username & password
	e.POST("/users", controllers.CreateUsersController)
	e.POST("/login", controllers.LoginUsersController)
	e.GET("/users", controllers.GetUsersController)

	// auth basic
	// auth := e.Group("basic")
	// auth.Use(middleware.BasicAuth(mid.AuthBasicMiddleware))
	// auth.GET("", controllers.GetUsersController)

	jwt := e.Group("jwt")
	jwt.Use(middleware.JWT([]byte(constants.KEY_JWT)))
	jwt.GET("", controllers.GetUsersController)

	e.GET("/users/companies", controllers.GetUsersCompanyController)
	e.POST("/users/companies", controllers.CreateUserCompaniesController)
	return e
}
