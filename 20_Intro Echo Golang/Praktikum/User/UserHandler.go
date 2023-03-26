package User

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type User struct {
	Id       int    `json:"id" form:"id"`
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

var users []User

// get all users
func GetUsersHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success get all users",
		"users":    users,
	})
}

// get user by id
func GetUserHandler(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	user := GetUserController(id)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success get user by id",
		"users":    user,
	})
}

// delete user by id
func DeleteUserHandler(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	DeleteUserController(id)
	return c.NoContent(http.StatusNoContent)
}

// update user by id
func UpdateUserHandler(c echo.Context) error {
	user := User{}
	c.Bind(&user)
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	UpdateUserController(user, id)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success update user",
		"user":     users[id-1],
	})
}

// create new user
func CreateUserHandler(c echo.Context) error {
	// binding data
	user := User{}
	c.Bind(&user)
	user = CreateUserController(user)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success create user",
		"user":     user,
	})
}
