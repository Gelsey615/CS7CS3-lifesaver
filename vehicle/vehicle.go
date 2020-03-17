package vehicle

import (
  "db"
  "googleMapAPI"
  "errors"
  "fmt"
  "encoding/json"
  "strings"
  "strconv"
)

const (
  ambulance = iota
  fireTruck
  policeCar
  bus
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
  if typ >= ambulance && typ < typCount {
    return nil
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
  var v Vehicle
  if err := json.Unmarshal([]byte(doc), &v); err != nil {
    return "", err
  }
  return fmt.Sprintf("{\"des_lat\":%f,\"des_ln\":%f}", v.DesLat, v.DesLong), nil
}

func RequestRoutePlan(id string) (string, error){
  doc, err := db.GetFromDB(db.ColVehicle, id)
  if err != nil {
    return "", err
  }
  var v Vehicle
  if err := json.Unmarshal([]byte(doc), &v); err != nil {
    return "", err
  }

  return googleMapAPI.GetRoute(v.Latitude, v.Longitude, v.DesLat, v.DesLong)
}

func DispatchVehicle(dispatchInfo map[int]int, disasterId string) error {
  for typ, num := range dispatchInfo {
    doc, err := db.QueryDB(db.ColVehicle, "{\"eq\": " + strconv.Itoa(typ) + ", \"in\": [\"type\"]}")
    if err != nil {
      return err
    }
    var vehicles map[string]Vehicle
    if err := json.Unmarshal([]byte(doc), &vehicles); err != nil {
      return err
    }
    err = selectVehicleForDispatch(vehicles, num, disasterId)
    if err != nil {
      return err
    }
  }
  return nil
}

func selectVehicleForDispatch(vehicles map[string]Vehicle, num int, disasterId string) error {
  selectedV := make([]string, 0)
  for vid, v := range vehicles {
    if v.DisasterId == "" {
      selectedV = append(selectedV, vid)
    }
    if len(selectedV) == num {
      break
    }
  }
  fmt.Println("disaster: "+disasterId+" selected vehicle: "+strings.Join(selectedV, ","))
  return db.UpdateToDB(db.ColVehicle, strings.Join(selectedV, ","), "{\"disaster_id\":\""+disasterId+"\"}")
}
