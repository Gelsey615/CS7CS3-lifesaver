package disaster

import (
  "fmt"
  "net/http"

  "github.com/webx-top/echo"
)

// type of disasters
const (

)

type Disaster struct {
  ID string `json:"Id"`
  Longitude float64 `json:"Longitude"`
  Latitude float64 `json:"Lat"`
  Radius float64 `json:"R"`
  Type string `json:"T"`
  Lvl int32 `json:"Lvl"`
  People int `json:"People"`
  StartTime int64 `json:"Start_time"`
  EndTime int64 `json:"End_time"`
}

func ReportDisaster(d Disaster) string {
  return d.Type
}

func GetDisaster(c echo.Context) error {
  // Bind the input data to ExampleRequest
  exampleRequest := new(Disaster)
  if err := c.Bind(exampleRequest); err != nil {
    return err
  }
  return c.JSONBlob([]byte(fmt.Sprintf(`{
        "longitude": %q,
        "altitude": %q,
        "r": "%q",
        "t": "%s"
      }`, exampleRequest.Longitude, exampleRequest.Latitude, exampleRequest.Radius, exampleRequest.Type)),http.StatusOK)
}
