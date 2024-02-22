package main

import (
  "fmt"
  "github.com/it3rno/balluhbac_encoder/handler"
  "github.com/labstack/echo/v4"
)

func main() {
  app := echo.New()

  homeHandler := handler.HomeHandler{}
  app.GET("/", homeHandler.HandleHome)

  app.start(":8080")
}
