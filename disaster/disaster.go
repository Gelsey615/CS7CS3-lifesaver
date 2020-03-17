package disaster

import (
  "encoding/json"
  "strconv"
  "time"
  "db"
  "googleMapAPI"
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
  return db.UpdateToDB(db.ColDisaster, id, "{\"end_time\":"+strconv.FormatInt(sec, 10)+"}")
}

func GetAllDisasters() (string, error) {
  str, err := db.QueryDB(db.ColDisaster, "{\"eq\": 0, \"in\": [\"end_time\"]}")
  return str, err
}
