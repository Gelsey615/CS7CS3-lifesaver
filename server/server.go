package main

import (
	"net/http"
	"github.com/webx-top/echo"
	"github.com/webx-top/echo/engine/standard"
  "disaster"
  "encoding/json"
	"fmt"
)

func main() {
	e := echo.New()
	e.Get("/", func(c echo.Context) error {
		return c.String("Hello, World!", http.StatusOK)
	})

	e.Get("/disaster/getall", getAllDisasters)
	e.Post("/disaster/report", reportDisaster)

	e.Run(standard.New(":1323"))
}

func reportDisaster(c echo.Context) error {
	d := disaster.Disaster{}
	if err := c.Bind(&d); err != nil {
			return err
	}
	return c.String(d.Type, http.StatusOK)
	/*disasterID := disaster.ReportDisaster(d)
	if disasterID == "" {
		return c.String("", http.StatusInternalServerError)
	} else {
		return c.String(disasterID, http.StatusOK)
	}*/
}

func getAllDisasters(c echo.Context) error {
  d := &disaster.Disaster{
    Longitude: 1,
    Latitude: 1,
    Radius: 1,
    Type: "fire",
  }
  b, err := json.Marshal(d)
    if err != nil {
        fmt.Println(err)
        return c.String("Unable to marshal", http.StatusOK)
    }
  return c.JSONBlob([]byte(string(b)),http.StatusOK)
}
