package main

import (
  "vehicle"
  "db"
  "github.com/webx-top/echo"
  "net/http"
  "fmt"
)

func RegisterNewVehicle(c echo.Context) error {
  v := new(db.Vehicle)
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

func RequestRoutePlan(c echo.Context) error {
  route, err := vehicle.RequestRoutePlan(c.Form("id"))
  if err != nil {
		fmt.Println(err.Error())
		return c.String(fmt.Sprintf("request route failed %s", err.Error()), http.StatusInternalServerError)
	} else {
		return c.JSONBlob([]byte(route),http.StatusOK)
	}
}

func DispatchVehicle(c echo.Context) error {
  info := make(map[int]int)
  if err := c.MustBind(&info); err != nil {
		return c.String(err.Error(), http.StatusInternalServerError)
	}

  err := vehicle.DispatchVehicle(info, c.Form("disaster_id"))
  if err != nil {
		fmt.Println(err.Error())
		return c.String("dispatch vehicle failed", http.StatusInternalServerError)
	} else {
		return c.String("fine", http.StatusOK)
	}
}
