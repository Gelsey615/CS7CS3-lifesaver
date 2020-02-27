package main

import (
	"net/http"
	"github.com/webx-top/echo"
	"github.com/webx-top/echo/engine/standard"
	"fmt"
	"io/ioutil"
	"bytes"
	"strings"
	"db"
)

// db connection, collections
const (
	DB = "http://localhost:8080/"
	ColDisaster = "disaster"
	ColVehicle = "vehicle"
)

func main() {
	err := checkDBCollection()
	if err != nil {
		panic(err.Error())
	}
	e := echo.New()

	e.Get("/disaster/getall", GetAllDisasters)
	e.Post("/disaster/report", ReportDisaster)
	e.Post("/disaster/finish", FinishDisaster)

	e.Run(standard.New(":1323"))
}

func checkDBCollection() error {
	// Get all collections in DB
	resp, err := http.Get(DB+"all")
	defer resp.Body.Close()
	if err != nil {
		fmt.Println("DB Error: get all collections failed, ", err.Error())
		return err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("DB Error: read resp err, ", err.Error())
		return err
	}

	// Parse collections
	byteReader := bytes.NewReader(body)
  byteStorage := make([]byte, 1028)
  n, err := byteReader.ReadAt(byteStorage, 2)
	if err.Error() != "EOF" {
		fmt.Println("DB Error: collection parse err, ", err.Error())
	}
	cols := []string{}
	if n > 2 {
		cols = strings.Split(string(byteStorage[:n-2]), "\",\"")
	}
	// Compare collections
	collection := []string{ColDisaster, ColVehicle}
	for _, c := range collection {
		// create necessary collection
		if !contain(cols, c) {
			err = db.CreateCollection(c)
			if err != nil {
				fmt.Printf("DB Error: create collections %s failed, %s\n", c, err.Error())
				return err
			}
			err = db.CreateIndex(c)
			if err != nil {
				fmt.Printf("DB Error: collections %s create index failed, %s\n", c, err.Error())
				return err
			}
		}
	}
	return nil
}

func contain(arr []string, str string) bool {
	for _, s := range arr {
		if s == str {
			return true
		}
	}
	return false
}
