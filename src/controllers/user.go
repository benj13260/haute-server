package controllers

import (
	"ben/haute/db"
	"ben/haute/models"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func CreateUser(c echo.Context) error {
	u := &models.User{}
	if err := c.Bind(u); err != nil {
		return err
	}
	result := db.DB.Create(&u)
	if result.Error != nil {
		return c.JSON(http.StatusNotFound, result.Error)
	}
	return c.JSON(http.StatusOK, u)
}

func GetUser(c echo.Context) error {
	id := c.Param("id")
	u := &models.User{}
	result := db.DB.First(u, "id = ?", id)
	if result.Error != nil {
		return c.JSON(http.StatusNotFound, result.Error)
	}
	return c.JSON(http.StatusOK, u)

}

func UpdateUser(c echo.Context) error {
	u := &models.User{}
	if err := c.Bind(u); err != nil {
		return err
	}
	result := db.DB.Updates(&u)
	if result.Error != nil {
		return c.JSON(http.StatusNotFound, result.Error)
	}
	return c.JSON(http.StatusOK, u)
}

func DeleteUser(c echo.Context) error {
	ids := c.Param("id")
	id, err := uuid.Parse(ids)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Incorrect ID")
	}
	u := &models.User{ID: id}
	result := db.DB.Delete(u)
	if result.Error != nil {
		return c.JSON(http.StatusNotFound, result.Error)
	}
	return c.JSON(http.StatusOK, u)
}

func GetAllUsers(c echo.Context) error {
	var users []models.User
	result := db.DB.Find(&users)
	if result.Error != nil {
		return c.JSON(http.StatusNotFound, result.Error)
	}
	return c.JSON(http.StatusOK, users)
}

func SearchUsers(c echo.Context) error {
	var users []models.User
	var result *gorm.DB
	db.DB.Model(&models.User{}).Select("name like ?").Where("name LIKE ?", "group%").Group("name").First(&result)

	if result.Error != nil {
		return c.JSON(http.StatusNotFound, result.Error)
	}
	return c.JSON(http.StatusOK, users)
}

func GetUserTypes(c echo.Context) error {
	var userType []models.UserType
	result := db.DB.Find(&userType)
	if result.Error != nil {
		return c.JSON(http.StatusNotFound, result.Error)
	}
	return c.JSON(http.StatusOK, userType)
}
