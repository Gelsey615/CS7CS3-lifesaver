package main

import (
  "user"
  "db"
	"github.com/webx-top/echo"
  "net/http"
  "fmt"
)

func RegisterNewUser(c echo.Context) error {
  u := new(db.User)
  if err := c.MustBind(u); err != nil {
		return c.String("binding err", http.StatusInternalServerError)
	}
  err :=  user.RegisterNewUser(u)
  if err != nil {
    return c.String(fmt.Sprintf("register failed: %s", err.Error()), http.StatusInternalServerError)
  } else {
    return c.NoContent(http.StatusOK)
  }
}

func Login(c echo.Context) error {
  u := new(db.User)
  if err := c.MustBind(u); err != nil {
		return c.String("binding err", http.StatusInternalServerError)
	}
  err :=  user.Login(u)
  if err != nil {
    return c.String(fmt.Sprintf("login failed: %s", err.Error()), http.StatusInternalServerError)
  } else {
    return c.NoContent(http.StatusOK)
  }
}
