package disaster

import (
  "fmt"
  "net/http"

  "github.com/webx-top/echo"
)

type Disaster struct {
  Longitude float32 `json:"longitude"`
  Altitude float32 `json:"altitude"`
  Radial float32 `json:"r"`
  Type string `json:"t"`
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
        "r": "%s",
        "t": "%s"
      }`, exampleRequest.Longitude, exampleRequest.Altitude, exampleRequest.Radial, exampleRequest.Type)),http.StatusOK)
}
