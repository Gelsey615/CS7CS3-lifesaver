package main

import (
  "disaster"
  "db"
	"github.com/webx-top/echo"
  "net/http"
  "fmt"
  "strconv"
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

func SetAssemblyPoint(c echo.Context) error {
  info := make(map[string]string)
  if err := c.MustBind(&info); err != nil {
		return c.String(err.Error(), http.StatusInternalServerError)
	}
  latFloat, err := strconv.ParseFloat(info["assembly_lat"], 64)
  if err != nil {
    return c.String("invalid assembly lat", http.StatusInternalServerError)
  }
  lnFloat, err := strconv.ParseFloat(info["assembly_ln"], 64)
  if err != nil {
    return c.String("invalid assembly ln", http.StatusInternalServerError)
  }
  err = disaster.SetAssemblyPoint(info["id"], latFloat, lnFloat)
  if err != nil {
    return c.String(err.Error(), http.StatusInternalServerError)
  } else {
    return c.NoContent(http.StatusOK)
  }
}

func FinishDisaster(c echo.Context) error {
  info := make(map[string]string)
  if err := c.MustBind(&info); err != nil {
		return c.String(err.Error(), http.StatusInternalServerError)
	}
  err := disaster.FinishDisaster(info["id"])
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
  info := make(map[string]string)
  if err := c.MustBind(&info); err != nil {
		return c.String(err.Error(), http.StatusInternalServerError)
	}
  vehicles, err := disaster.QueryDisasterReqVehicles(info["id"])
  if err != nil {
		fmt.Println(err.Error())
		return c.String("get all disasters failed", http.StatusInternalServerError)
	} else {
		return c.JSONBlob([]byte(vehicles),http.StatusOK)
	}
}
