package controllers

import (
	"github.com/labstack/echo"
)

type ResponseError struct {
	Error string `json:"error"`
}

func RespondWithError(c echo.Context, status int, err error) error {
	return c.JSON(status, ResponseError{Error: err.Error()})
}
