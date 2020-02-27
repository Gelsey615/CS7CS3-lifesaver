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

func UpdateToDB(col string, id string, doc string) error {
  resp, err := http.PostForm(DB+Merge, url.Values{"col": {col}, "doc": {doc}, "id": {id}})
  defer resp.Body.Close()
  return err
}

func GetFromDB(col string, query string) (string, error) {
  resp, err := http.PostForm(DB+Query, url.Values{"col": {col}, "q": {query}})
  defer resp.Body.Close()
  body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
  return string(body), err
}
