package main

import (
  "disaster"
  "db"
	"github.com/webx-top/echo"
  "net/http"
  "fmt"
)

func ReportDisaster(c echo.Context) error {
	d := new(db.Disaster)
	if err := c.MustBind(d); err != nil {
		return c.String("binding err", http.StatusInternalServerError)
	}
	disasterID, err := disaster.ReportDisaster(d)
	if err != nil {
		fmt.Println(err.Error())
		return c.String("report failed", http.StatusInternalServerError)
	} else {
		return c.String(disasterID, http.StatusOK)
	}
}

func FinishDisaster(c echo.Context) error {
  id := c.Form("id")
  err := disaster.FinishDisaster(id)
  if err != nil {
		fmt.Println(err.Error())
		return c.String("finish failed", http.StatusInternalServerError)
	} else {
		return c.NoContent(http.StatusOK)
	}
}

func GetAllDisasters(c echo.Context) error {
  disasters, err := disaster.GetAllDisasters()
  if err != nil {
		fmt.Println(err.Error())
		return c.String("get all disasters failed", http.StatusInternalServerError)
	} else {
		return c.JSONBlob([]byte(disasters),http.StatusOK)
	}
}

func QueryDisasterReqVehicles(c echo.Context) error {
  id := c.Form("id")
  vehicles, err := disaster.QueryDisasterReqVehicles(id)
  if err != nil {
		fmt.Println(err.Error())
		return c.String("get all disasters failed", http.StatusInternalServerError)
	} else {
		return c.JSONBlob([]byte(vehicles),http.StatusOK)
	}
}
