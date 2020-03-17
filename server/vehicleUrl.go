package main

import (
  "vehicle"
  "github.com/webx-top/echo"
  "net/http"
  "fmt"
)

func RegisterNewVehicle(c echo.Context) error {
  v := new(vehicle.Vehicle)
  if err := c.MustBind(v); err != nil {
		return c.String("binding err", http.StatusInternalServerError)
	}
  vehicleId, err := vehicle.RegisterNewVehicle(v)
  if err != nil {
    fmt.Println(err.Error())
    return c.String("register failed", http.StatusInternalServerError)
  } else {
    return c.String(vehicleId, http.StatusOK)
  }
}

func QueryVehicles(c echo.Context) error {
  info, err := vehicle.QueryVehicles()
  if err != nil {
    fmt.Println(err.Error())
    return c.String("query failed", http.StatusInternalServerError)
  } else {
    return c.JSONBlob([]byte(info),http.StatusOK)
  }
}

func UpdateVehicle(c echo.Context) error {
  info , err := vehicle.UpdateVehicle(c.Form("id"), c.Form("lat"), c.Form("ln"))
  if err != nil {
		fmt.Println(err.Error())
		return c.String("update vehicle failed", http.StatusInternalServerError)
	} else {
		return c.JSONBlob([]byte(info),http.StatusOK)
	}
}
