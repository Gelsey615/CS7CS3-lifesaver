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

	e.Run(standard.New(":1323"))
}

func getAllDisasters(c echo.Context) error {
  d := &disaster.Disaster{
    Longitude: 1,
    Altitude:1,
    Radial:1,
    Type:"fire",
  }
  b, err := json.Marshal(d)
    if err != nil {
        fmt.Println(err)
        return c.String("Unable to marshal", http.StatusOK)
    }
  return c.JSONBlob([]byte(string(b)),http.StatusOK)
}
