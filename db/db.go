package db

import (
  "net/http"
  "net/url"
  "io/ioutil"
  "fmt"

)

const (
  DB = "http://localhost:8080/"
  Insert = "insert"
  Merge = "merge"
  Query = "query"
  Get = "get"
	ColDisaster = "disaster"
  ColVehicle = "vehicle"
)

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
