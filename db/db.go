package db

import (
  "net/http"
  "net/url"
  "io/ioutil"
  "fmt"

)

const (
  DB = "http://rdoc.zhaojiajin.club:80/"
  Insert = "insert"
  Merge = "merge"
  Query = "query"
  Get = "get"
	ColDisaster = "disaster"
  ColVehicle = "vehicle"
)

type Vehicle struct {
  Type int `json:"type"`
  Longitude float64 `json:"longitude"`
  Latitude float64 `json:"latitude"`
  DisasterId string `json:"disaster_id"`
  DesLong float64 `json:"des_ln"`
  DesLat float64 `json:"des_lat"`
}

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

func CreateCollection(c string) error {
	fmt.Printf("Creating DB collection %s\n", c)
	resp, err := http.Get(DB+"create?col="+c)
	defer resp.Body.Close()
	return err
}

func CreateIndex(c string) error {
	switch c {
	case ColDisaster:
		fmt.Printf("Creating DB collection %s index \"end_time\"\n", c)
		resp, err := http.PostForm(DB+"index", url.Values{"col": {ColDisaster}, "path": {"end_time"}})
	  defer resp.Body.Close()
	  return err
  case ColVehicle:
    fmt.Printf("Creating DB collection %s index \"disaster_id\", \"type\"\n", c)
    resp, err := http.PostForm(DB+"index", url.Values{"col": {ColVehicle}, "path": {"disaster_id"}})
	  defer resp.Body.Close()
	  if err != nil {
      return err
    }
    resp2, err := http.PostForm(DB+"index", url.Values{"col": {ColVehicle}, "path": {"type"}})
	  defer resp2.Body.Close()
    return err
	}
	return nil
}

func InsertToDB(col string, doc string) (string, error) {
  resp, err := http.PostForm(DB+Insert, url.Values{"col": {col}, "doc": {doc}})
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

func GetFromDB(col string, id string) (string, error) {
  resp, err := http.PostForm(DB+Get, url.Values{"col": {col}, "id": {id}})
  defer resp.Body.Close()
  if err != nil {
    return "", err
  }
  body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
    return "", err
  }
  return string(body), nil
}

func UpdateToDB(col string, id string, doc string) error {
  resp, err := http.PostForm(DB+Merge, url.Values{"col": {col}, "doc": {doc}, "id": {id}})
  defer resp.Body.Close()
  return err
}

func QueryDB(col string, query string) (string, error) {
  resp, err := http.PostForm(DB+Query, url.Values{"col": {col}, "q": {query}})
  defer resp.Body.Close()
  body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
  return string(body), err
}
