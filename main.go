package main

import (
  "github.com/labstack/echo"
  "github.com/labstack/echo/engine/standard"
  "github.com/labstack/echo/middleware"
  "github.com/matthewharwood/morningharwood/api"
)

const (
  dev  string = "client/dev"
)

func main() {
  e := echo.New()

  // Middleware
  e.Use(middleware.Logger())
  e.Use(middleware.Recover())


  // Route => handler
  e.Pre(middleware.RemoveTrailingSlash())
  e.Use(middleware.Static(dev))

  todoroutes.Init(e)

  // Start server
  e.Run(standard.New(":1323"))
}

