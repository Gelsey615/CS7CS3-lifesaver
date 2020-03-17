package vehicle

import (
  "db"
  "errors"
  "fmt"
  "encoding/json"
)

const (
  ambulance = iota
  fireTruck
  policeCar
  typCount
)

type Vehicle struct {
  Type int `json:"type"`
  Longitude float64 `json:"longitude"`
  Latitude float64 `json:"latitude"`
  DisasterId string `json:"disaster_id"`
  DesLong float64 `json:"des_ln"`
  DesLat float64 `json:"des_lat"`
}

func typeCheck(typ int) error {
  typs := []int{ambulance, fireTruck, policeCar}
  for _, t := range typs {
    if t == typ {
      return nil
    }
  }
  return errors.New(fmt.Sprintf("invalid type %i", typ))
}

func RegisterNewVehicle(v *Vehicle) (string, error) {
  //type check
  err := typeCheck(v.Type)
  if err != nil {
    return "", err
  }
  //save to db
  doc, err := json.Marshal(v)
  if err != nil {
    return "", err
  }
  id, err := db.InsertToDB(db.ColVehicle, string(doc))
  return id, err
}

func QueryVehicles() (string, error){
  doc, err := db.QueryDB(db.ColVehicle, "{\"eq\": \"\", \"in\": [\"disaster_id\"]}")
  if err != nil {
    return "", err
  }
  var dat map[string]Vehicle
  if err := json.Unmarshal([]byte(doc), &dat); err != nil {
    return "", err
  }

  count := make(map[int]int)
  for _, v := range dat {
    count[v.Type]++
  }

  str, err := json.Marshal(count)
  if err != nil {
    return "", err
  }
  return string(str), err
}

func UpdateVehicle(id string, lat string, ln string) (string, error) {
  err := db.UpdateToDB(db.ColVehicle, id, "{\"latitude\":"+lat+",\"longitude\":"+ln+"}")
  if err != nil {
    return "", err
  }
  doc, err := db.GetFromDB(db.ColVehicle, id)
  if err != nil {
    return "", err
  }
  var dat Vehicle
  if err := json.Unmarshal([]byte(doc), &dat); err != nil {
    return "", err
  }
  return fmt.Sprintf("{\"des_lat\":%f,\"des_ln\":%f}", dat.DesLat, dat.DesLong), nil
}
