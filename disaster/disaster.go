package disaster

import (
  "net/http"
  "encoding/json"
  "io/ioutil"
  "net/url"
  "strconv"
  "time"
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
  id, err := saveToDB(string(doc))
  return id, err
}

func FinishDisaster(id string) error {
  now := time.Now()
  sec := now.Unix()
  return updateToDB(id, "{\"end_time\":"+strconv.FormatInt(sec, 10)+"}")
}

func GetAllDisasters() (string, error) {
  str, err := getFromDB()
  return str, err
}

func saveToDB(doc string) (string, error) {
  resp, err := http.PostForm(DB+Insert, url.Values{"col": {ColDisaster}, "doc": {doc}})
  defer resp.Body.Close()
  if err != nil {
    return "", err
  }
  id, err := ioutil.ReadAll(resp.Body)
	if err != nil {
    return "", err
  }
  return string(id), nil
}

func updateToDB(id string, doc string) error {
  resp, err := http.PostForm(DB+Merge, url.Values{"col": {ColDisaster}, "doc": {doc}, "id": {id}})
  defer resp.Body.Close()
  return err
}

func getFromDB() (string, error) {
  resp, err := http.PostForm(DB+Query, url.Values{"col": {ColDisaster}, "q": {"{\"eq\": 0, \"in\": [\"end_time\"]}"}})
  defer resp.Body.Close()
  body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
  return string(body), err
}
