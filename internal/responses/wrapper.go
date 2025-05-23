package responses

import "github.com/labstack/echo/v4"

type Error struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
}

type Data struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func Response(c echo.Context, code int, data any) error {
	return c.JSON(code, data)
}

func SuccessResponse(c echo.Context, code int, message string, data any) error {
	return Response(c, code, Data{
		Code:    code,
		Message: message,
		Data:    data,
	})
}

func ErrorResponse(c echo.Context, code int, message string) error {
	return Response(c, code, Error{
		Code:  code,
		Error: message,
	})
}
