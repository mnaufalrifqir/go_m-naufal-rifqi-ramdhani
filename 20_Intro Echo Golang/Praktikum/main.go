package main

import (
	"Praktikum/User"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	// routing with query parameter
	e.GET("/users", User.GetUsersHandler)
	e.POST("/users", User.CreateUserHandler)
	e.GET("/users/:id", User.GetUserHandler)
	e.PUT("/users/:id", User.UpdateUserHandler)
	e.DELETE("/users/:id", User.DeleteUserHandler)
	// start the server, and log if it fails
	e.Logger.Fatal(e.Start(":8000"))
}
