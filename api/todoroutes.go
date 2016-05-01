package todoroutes

import (
  "github.com/labstack/echo"
  "github.com/matthewharwood/morningharwood/api/controller"
)

func Init(e *echo.Echo) {
  e.Get("/api/todos", todocontroller.GetAll)
}