package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

// I have kept this outside , you may have in config file.
// This is required to be outside so that we can mock this url
var url = "http://ip.jsontest.com/"

func main() {

	var result, err = GetData()
	fmt.Println(result, err)
}

// GetData  => makes a http call to a dummy service and return data or error
func GetData() (m MyIP, err error) {

	resp, err := http.Get(url)

	if err == nil {
		if resp.StatusCode != 400 {

			body, err := ioutil.ReadAll(resp.Body)

			bytBodyErr := []byte(body)
			if err == nil {
				_ = json.Unmarshal(bytBodyErr, &m)
			}
		} else {
			return m, errors.New("400 Bad Request")
		}
	}
	return m, err
}

//MyIP struct for storing IP data
type MyIP struct {
	IP string `json:"ip"`
}
