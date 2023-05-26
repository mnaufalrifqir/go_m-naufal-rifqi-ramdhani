package route

import (
	"code_competence/constants"
	"code_competence/controller"
	"code_competence/middlewares"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func StartRoute() *echo.Echo {
	e := echo.New()

	middlewares.LogMiddleware(e)

	noAuth := e.Group("")
	noAuth.POST("/register", controller.Register)
	noAuth.POST("/login", controller.Login)

	Auth := e.Group("/auth", echojwt.JWT([]byte(constants.SECRET_JWT)))
	Auth.GET("/items", controller.GetItemsController)
	Auth.GET("/items/:id", controller.GetItemController)
	Auth.POST("/items", controller.CreateItemController)
	Auth.PUT("/items/:id", controller.UpdateItemController)
	Auth.DELETE("/items/:id", controller.DeleteItemController)
	Auth.GET("/items/category/:category_id", controller.GetItemsByCategoryIdController)
	Auth.GET("/items", controller.GetItemByNameController)
	return e
}
