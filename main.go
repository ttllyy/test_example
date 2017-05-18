package main

import (
	"net/http"
	//"strings"
	"fmt"
	"io/ioutil"
)

type RestClient struct {
	Url        string       // The location of the server like "http://localhost:8065"
	HttpClient *http.Client // The http client
}


//var Rclient *RestClient

func NewClient() *RestClient {
	return &RestClient{"http://127.0.0.1:8081/", &http.Client{}}
}

func (c *RestClient) DoGet() (*http.Response, error) {
	rq, _ := http.NewRequest("GET", c.Url, nil)
	fmt.Println(rq.URL)
	return c.HttpClient.Do(rq)
}


func main() {
	client := NewClient()
	res, err := client.DoGet()
	if err != nil {
		fmt.Println(err.Error())
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	fmt.Println(string(body))


	fmt.Println("exit")
}
