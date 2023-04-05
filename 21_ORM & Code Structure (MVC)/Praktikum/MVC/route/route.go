package route

import (
	"mvc_alterra/constants"
	"mvc_alterra/controller"
	"mvc_alterra/middlewares"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func StartRoute() *echo.Echo {
	e := echo.New()

	// Section 22
	middlewares.LogMiddleware(e)

	e.POST("/user", controller.CreateUserController)
	e.POST("/user/login", controller.LoginUserController)

	u := e.Group("/users")
	u.Use(echojwt.JWT([]byte(constants.SECRET_JWT)))
	u.GET("", controller.GetUsersController)
	u.GET("/:id", controller.GetUserController)
	u.DELETE("/:id", controller.DeleteUserController)
	u.PUT("/:id", controller.UpdateUserController)

	b := e.Group("/books")
	b.Use(echojwt.JWT([]byte(constants.SECRET_JWT)))
	b.GET("", controller.GetBooksController)
	b.GET("/:id", controller.GetBookController)
	b.POST("", controller.CreateBookController)
	b.DELETE("/:id", controller.DeleteBookController)
	b.PUT("/:id", controller.UpdateBookController)

	// Section 21
	// e.GET("/users", controller.GetUsersController)
	// e.GET("/users/:id", controller.GetUserController)
	// e.POST("/users", controller.CreateUserController)
	// e.DELETE("/users/:id", controller.DeleteUserController)
	// e.PUT("/users/:id", controller.UpdateUserController)

	// e.GET("/books", controller.GetBooksController)
	// e.GET("/books/:id", controller.GetBookController)
	// e.POST("/books", controller.CreateBookController)
	// e.DELETE("/books/:id", controller.DeleteBookController)
	// e.PUT("/books/:id", controller.UpdateBookController)

	// e.GET("/blogs", controller.GetBlogsController)
	// e.GET("/blogs/:id", controller.GetBlogController)
	// e.POST("/blogs", controller.CreateBlogController)
	// e.DELETE("/blogs/:id", controller.DeleteBlogController)
	// e.PUT("/blogs/:id", controller.UpdateBlogController)

	return e
}
