package controllers

import (
	"github.com/labstack/echo"
	"net/http"
	"strconv"
	"synack/models"
	"synack/services"
)

func GetUser(c echo.Context) error {

	userID := c.Param("id")
	service := services.UserService{}

	intUserID, err := strconv.Atoi(userID)
	if err != nil {
		return RespondWithError(c, http.StatusBadRequest, err)
	}

	user, err := service.GetUserByID(intUserID)
	if err != nil {
		return RespondWithError(c, http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, user)
}

func SaveUser(c echo.Context) error {

	userReq := new(models.User)
	if err := c.Bind(userReq); err != nil {
		return RespondWithError(c, http.StatusBadRequest, err)
	}

	service := services.UserService{}
	user, err := service.SaveUser(userReq)
	if err != nil {
		return RespondWithError(c, http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, user)
}

func UpdateUser(c echo.Context) error {

	userID := c.Param("id")
	service := services.UserService{}

	intUserID, err := strconv.Atoi(userID)
	if err != nil {
		return RespondWithError(c, http.StatusBadRequest, err)
	}

	userReq := new(models.User)
	if err = c.Bind(userReq); err != nil {
		return RespondWithError(c, http.StatusBadRequest, err)
	}

	user, err := service.UpdateUser(intUserID, userReq)
	if err != nil {
		return RespondWithError(c, http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, user)
}

func DeleteUser(c echo.Context) error {

	userID := c.Param("id")
	service := services.UserService{}

	intUserID, err := strconv.Atoi(userID)
	if err != nil {
		return RespondWithError(c, http.StatusBadRequest, err)
	}

	user, err := service.DeleteUser(intUserID)
	if err != nil {
		return RespondWithError(c, http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, user)
}
