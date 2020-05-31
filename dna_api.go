package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type FullResponse struct {
	Response Response `json:"response"`
}

type Response struct {
	Docs []Doc `json:"docs"`
}

type Doc struct {
	Abstract []string `json:"abstract"`
}

func main() {
	url := "http://api.plos.org/search?q=title:DNA"
	//url := "http://dummy.restapiexample.com/api/v1/employees"
	//url := "http://api.open-notify.org/astros.json"

	resp, _ := http.Get(url)
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	var data FullResponse
	json.Unmarshal(body, &data)
}
