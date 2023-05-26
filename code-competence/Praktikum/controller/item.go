package controller

import (
	"code_competence/database"
	"code_competence/model"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// Melihat Item
func GetItemsController(c echo.Context) error {
	var items []model.Item

	if err := database.DB.Find(&items).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all items",
		"users":   items,
	})
}

// Melihat Item berdasarkan ID
func GetItemController(c echo.Context) error {
	item := model.Item{}
	ItemId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	if err1 := database.DB.First(&item, ItemId).Error; err1 != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err1.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get item by id",
		"user":    item,
	})
}

// Nambah Item
func CreateItemController(c echo.Context) error {
	item := model.Item{}
	c.Bind(&item)

	if err := database.DB.Save(&item).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new item",
		"item":    item,
	})
}

// Update Item
func UpdateItemController(c echo.Context) error {
	item := model.Item{}
	ItemID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	if err := database.DB.First(&item, ItemID).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	c.Bind(&item)
	if err := database.DB.Save(&item).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]string{
		"message": "item updated successfully",
	})
}

// Menghapus Item
func DeleteItemController(c echo.Context) error {
	item := model.Item{}
	ItemId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	if err := database.DB.Delete(&item, ItemId).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "Item deleted successfully",
	})
}

func GetItemsByCategoryIdController(c echo.Context) error {
	var items []model.Item
	categoryId, err := strconv.Atoi(c.Param("category_id"))
	if err != nil {
		return err
	}
	if err := database.DB.Where("`category_id` = ?", categoryId).Find(&items).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get items by category id",
		"item":    items,
	})
}

// Search Item Berdasarkan Nama
func GetItemByNameController(c echo.Context) error {
	var items []model.Item
	name := c.QueryParam("keyword")

	if err := database.DB.Where("name LIKE ?", "%"+name+"%").Find(&items).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get item by name",
		"item":    items,
	})
}
