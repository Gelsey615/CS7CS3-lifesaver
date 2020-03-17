package disaster

import (
  "encoding/json"
  "strconv"
  "time"
  "db"
  "googleMapAPI"
)

type Disaster struct {
  Longitude float64 `json:"longitude"`
  Latitude float64 `json:"latitude"`
  Radius float64 `json:"radius"`
  Type string `json:"type"`
  Lvl int32 `json:"lvl"`
  People int `json:"people"`
  StartTime int64 `json:"start_time"`
  EndTime int64 `json:"end_time"`
  AssemblyLn float64 `json:"assembly_ln"`
  AssemblyLat float64 `json:"assembly_lat"`
  ReqRoute string `json:"req_route"`
}

func ReportDisaster(d *Disaster) (string, error) {
  route, err := googleMapAPI.GetRoute(d.Latitude, d.Longitude, d.AssemblyLat, d.AssemblyLn)
  if err != nil {
    return "", err
  }
  d.ReqRoute = route
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
  return db.UpdateToDB(db.ColDisaster, id, "{\"end_time\":"+strconv.FormatInt(sec, 10)+"}")
}

func GetAllDisasters() (string, error) {
  str, err := db.QueryDB(db.ColDisaster, "{\"eq\": 0, \"in\": [\"end_time\"]}")
  return str, err
}
