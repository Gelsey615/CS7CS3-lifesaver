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

func typeCheck(typ int) error {
  if typ >= ambulance && typ < typCount {
    return nil
  }
  return errors.New(fmt.Sprintf("invalid type %i", typ))
}

func RegisterNewVehicle(v *db.Vehicle) (string, error) {
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
  var dat map[string]db.Vehicle
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
  var v map[string]db.Vehicle
  if err := json.Unmarshal([]byte(doc), &v); err != nil {
    return "", err
  }
  return fmt.Sprintf("{\"des_lat\":%f,\"des_ln\":%f}", v[id].DesLat, v[id].DesLong), nil
}

func RequestRoutePlan(id string) (string, error){
  doc, err := db.GetFromDB(db.ColVehicle, id)
  if err != nil {
    return "", err
  }
  var v map[string]db.Vehicle
  if err := json.Unmarshal([]byte(doc), &v); err != nil {
    return "", err
  }
  if v[id].DisasterId == "" {
    return "", errors.New(fmt.Sprintf("vehicle %s is not dispatched", id))
  }

  myRoute, err := googleMapAPI.GetRoute(v[id].Latitude, v[id].Longitude, v[id].DesLat, v[id].DesLong)
  if err != nil {
    return "", err
  }


  doc, err = db.GetFromDB(db.ColDisaster, v[id].DisasterId)
  if err != nil {
    return "", err
  }
  var d map[string]db.Disaster
  if err := json.Unmarshal([]byte(doc), &d); err != nil {
    return "", err
  }

  return fmt.Sprintf("{\"req_route\":%s,\"my_route\":%s}", d[v[id].DisasterId].ReqRoute, myRoute), nil
}

func DispatchVehicle(dispatchInfo map[int]int, disasterId string) error {
  doc, err := db.GetFromDB(db.ColDisaster, disasterId)
  if err != nil {
    return err
  }

  var d map[string]db.Disaster
  if err := json.Unmarshal([]byte(doc), &d); err != nil {
    return err
  }

  for typ, num := range dispatchInfo {
    doc, err := db.QueryDB(db.ColVehicle, "{\"eq\": " + strconv.Itoa(typ) + ", \"in\": [\"type\"]}")
    if err != nil {
      return err
    }
    var vehicles map[string]db.Vehicle
    if err := json.Unmarshal([]byte(doc), &vehicles); err != nil {
      return err
    }
    err = selectAndDispatchVehicle(vehicles, num, disasterId, d[disasterId].Latitude, d[disasterId].Longitude)
    if err != nil {
      return err
    }
  }
  return nil
}

func selectAndDispatchVehicle(vehicles map[string]db.Vehicle, num int, disasterId string, lat float64, ln float64) error {
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
  queryStr := "{\"disaster_id\":\""+disasterId+"\",\"des_lat\":"+fmt.Sprintf("%f", lat)+",\"des_ln\":"+fmt.Sprintf("%f", ln)+"}"
  return db.UpdateToDB(db.ColVehicle, strings.Join(selectedV, ","), queryStr)
}
