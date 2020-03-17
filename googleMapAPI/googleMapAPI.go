package googleMapAPI

import (
  "googlemaps.github.io/maps"
  "github.com/kr/pretty"
  "context"
  "fmt"
  "encoding/json"
  "errors"
)

const (
  googleApiKey = "AIzaSyDzxCwIwkSWHHoO0Kj_SS2FSfG9Ebkt-B8"
)

func GetRoute(originLat float64, originLn float64, desLat float64, desLn float64) (string, error) {
  c, err := maps.NewClient(maps.WithAPIKey(googleApiKey))
	if err != nil {
    return "", errors.New(fmt.Sprintf("API connection error: %s", err))
	}
  fmt.Printf("%f,%f\n", originLat, originLn)
  fmt.Printf("%f,%f\n", desLat, desLn)
	r := &maps.DirectionsRequest{
		Origin:      fmt.Sprintf("%f,%f", originLat, originLn),
		Destination: fmt.Sprintf("%f,%f", desLat, desLn),
	}
	route, _, err := c.Directions(context.Background(), r)
	if err != nil || len(route) <= 0 {
    return "", errors.New(fmt.Sprintf("get route error: %s", err))
	}
	pretty.Println(route[0])
  doc, err := json.Marshal(route[0])
  if err != nil {
    return "", err
  }
  pretty.Println(string(doc))
  return string(doc), nil
}
