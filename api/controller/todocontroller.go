package todocontroller

import (
  "github.com/labstack/echo"
  "net/http"
  "github.com/matthewharwood/morningharwood/api/dao"
  "fmt"
)



func GetAll(c echo.Context) error {
  ts, _ := tododao.All()
  fmt.Println(ts, "tttts")
  return c.JSON(http.StatusOK, ts)
}