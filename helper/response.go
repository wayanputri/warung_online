package helper

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Success(c echo.Context, message string, data any) error{
	return c.JSON(http.StatusOK, map[string]any{
		"status":"success",
		"message":message,
		"data":data,
	})
}

func SuccessWithOutData(c echo.Context, message string) error{
	return c.JSON(http.StatusOK, map[string]any{
		"status":"success",
		"message":message,
	})
}

func FailedNotFound(c echo.Context, message string, data any) error{
	return c.JSON(http.StatusNotFound, map[string]any{
		"status":"fail",
		"message": message,
		"data": data,
	})
}

func FailedRequest(c echo.Context,message string, data any) error{
	return c.JSON(http.StatusBadRequest, map[string]any{
		"status":"fail",
		"message": message,
		"data": data,
	})
}

func SuccessCreate(c echo.Context, message string, data any) error{
	return c.JSON(http.StatusCreated, map[string]any{
		"status":"success",
		"message": message,
		"data": data,
	})
}

func UnAutorization(c echo.Context, message string, data any)error{
	return c.JSON(http.StatusUnauthorized, map[string]any{
		"status":"fail",
		"message":message,
		"data":data,
	})
}

func Forbidden(c echo.Context, message string, data any) error{
	return c.JSON(http.StatusForbidden, map[string]any{
		"status":"fail",
		"message":message,
		"data": data,
	})
}

func InternalError(c echo.Context, message string, data any) error{
	return c.JSON(http.StatusInternalServerError, map[string]any{
		"status":"fail",
		"message":message,
		"data":data,
	})
}