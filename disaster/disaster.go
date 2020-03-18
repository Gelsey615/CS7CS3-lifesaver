package disaster

import (
  "encoding/json"
  "strconv"
  "time"
  "db"
  "googleMapAPI"
  "fmt"
  "strings"
)

func ReportDisaster(d *db.Disaster) (string, error) {
  route, err := googleMapAPI.GetRoute(d.Latitude, d.Longitude, d.AssemblyLat, d.AssemblyLn)
  if err != nil {
    return "", err
  }
  d.ReqRoute = route
  now := time.Now()
  sec := now.Unix()
  d.StartTime = sec
  doc, err := json.Marshal(d)
  if err != nil {
    return "", err
  }
  id, err := db.InsertToDB(db.ColDisaster, string(doc))
  return id, err
}

func FinishDisaster(id string) error {
  now := time.Now()
  sec := now.Unix()
  err := db.UpdateToDB(db.ColDisaster, id, "{\"end_time\":"+strconv.FormatInt(sec, 10)+"}")
  if err != nil {
    return err
  }

  doc, err := db.QueryDB(db.ColVehicle, "{\"eq\": \""+ id +"\", \"in\": [\"disaster_id\"]}")
  if err != nil {
    return err
  }
  var vehicles map[string]db.Vehicle
  if err := json.Unmarshal([]byte(doc), &vehicles); err != nil {
    return err
  }
  dispatchedV := make([]string, 0)
  for vid, _ := range vehicles {
    dispatchedV = append(dispatchedV, vid)
  }
  queryStr := "{\"disaster_id\":\"\",\"des_lat\":"+fmt.Sprintf("%f", 0.0)+",\"des_ln\":"+fmt.Sprintf("%f", 0.0)+"}"
  return db.UpdateToDB(db.ColVehicle, strings.Join(dispatchedV, ","), queryStr)
}

func GetAllDisasters() (string, error) {
  str, err := db.QueryDB(db.ColDisaster, "{\"eq\": 0, \"in\": [\"end_time\"]}")
  return str, err
}

func QueryDisasterReqVehicles(id string) (string, error) {
  vehicles, err := db.QueryDB(db.ColVehicle, "{\"eq\": \""+ id +"\", \"in\": [\"disaster_id\"]}")
  if err != nil {
    return "", err
  }
  return string(vehicles), nil
}
