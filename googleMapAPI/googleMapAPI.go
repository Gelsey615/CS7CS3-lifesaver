package googleMapAPI

import (
  "googlemaps.github.io/maps"
  //"github.com/kr/pretty"
  "context"
  "fmt"
  "encoding/json"
  "errors"
)

const (
  googleApiKey = "AIzaSyDzxCwIwkSWHHoO0Kj_SS2FSfG9Ebkt-B8"
)

type LocPoint struct {
  Lat float64 `json:"lat"`
  Lng float64 `json:"Lng"`
}

func GetRoute(originLat float64, originLn float64, desLat float64, desLn float64) (string, error) {
  c, err := maps.NewClient(maps.WithAPIKey(googleApiKey))
	if err != nil {
    return "", errors.New(fmt.Sprintf("API connection error: %s", err))
	}
  fmt.Printf("googleMapAPI origin: %f,%f\n", originLat, originLn)
  fmt.Printf("googleMapAPI des: %f,%f\n", desLat, desLn)
	r := &maps.DirectionsRequest{
		Origin:      fmt.Sprintf("%f,%f", originLat, originLn),
		Destination: fmt.Sprintf("%f,%f", desLat, desLn),
	}
	route, _, err := c.Directions(context.Background(), r)
	if err != nil || len(route) <= 0 {
    return "", errors.New(fmt.Sprintf("get route error: %s", err))
	}

  lc := []LocPoint{}
  for _, s := range route[0].Legs[0].Steps {
    locP := LocPoint{
      Lat: s.EndLocation.Lat,
      Lng: s.EndLocation.Lng,
    }
    lc = append(lc, locP)
  }
  doc, err := json.Marshal(lc)
  if err != nil {
    return "", err
  }

  return string(doc), nil
}
