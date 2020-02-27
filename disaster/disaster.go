package disaster

import (
  "encoding/json"
  "strconv"
  "time"
  "db"
)

// type of disasters
const (
  DB = "http://localhost:8080/"
  Insert = "insert"
  Merge = "merge"
  Query = "query"
	ColDisaster = "disaster"
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
}

func ReportDisaster(d *Disaster) (string, error) {
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
  str, err := db.GetFromDB(db.ColDisaster, "{\"eq\": 0, \"in\": [\"end_time\"]}")
  return str, err
}
