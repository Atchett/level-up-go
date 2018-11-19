package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// Config is the struct for reading the JSON into
type Config struct {
	Name     string `json:"SiteName"`
	URL      string `json:"SiteURL"`
	Database struct {
		Name     string
		Host     string
		Port     int
		Username string
		Password string
	}
}

func main() {
	// setup an empty instance of the Config struct
	conf := Config{}
	// read the file into the data variable
	data, err := ioutil.ReadFile("config.json")
	if err != nil {
		panic(err)
	}

	// unmarshall the file contents (data) directly into the conf variable (pointer)
	err = json.Unmarshal(data, &conf)
	if err != nil {
		panic(err)
	}

	// print the site details
	fmt.Printf("Site: %s (%s)\n", conf.Name, conf.URL)

	db := conf.Database

	//Print out database connection string
	// %s - string
	// %d - decimal (base 10)
	fmt.Printf(
		"DB: mysql://%s:%s@%s:%d/%s\n",
		db.Username,
		db.Password,
		db.Host,
		db.Port,
		db.Name,
	)
}
