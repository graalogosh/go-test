package handlers

import "github.com/labstack/echo/v4"

type rootResponse struct {
	Hello string `json:"hello"`
	World string `json:"world"`
}

func RootHandler(context echo.Context) error {
	context.JSON(200, rootResponse{
		Hello: "value1",
		World: "value2",
	})
	return nil
}
