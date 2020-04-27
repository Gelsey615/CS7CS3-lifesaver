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

func SetAssemblyPoint(id string, assemblyLat float64, assemblyLn float64 ) error {
  doc, err := db.GetFromDB(db.ColDisaster, id)
  if err != nil {
    return err
  }
  var d map[string]db.Disaster
  if err := json.Unmarshal([]byte(doc), &d); err != nil {
    return err
  }
  route, err := googleMapAPI.GetRoute(d[id].Latitude, d[id].Longitude, assemblyLat, assemblyLn)
  if err != nil {
    return err
  }
  updateStr := fmt.Sprintf("{\"req_route\":%s, \"assembly_lat\":%f, \"assembly_ln\":%f}", route, assemblyLat, assemblyLn)
  return db.UpdateToDB(db.ColDisaster, id, updateStr)
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
  if doc == "{}" {
    return nil
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
