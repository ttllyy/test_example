package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	"io/ioutil"
)

type Msg struct {
	Name	string		`json:"name"`
	Age		int			`json:"age"`
	Score	int			`json:"score"`
}

var res []byte

func handler(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	var m Msg
	json.Unmarshal(body, &m)
	m.Age++
	m.Score++
	o, _ := json.Marshal(m)
	w.Write(o)
}


func main() {
	fmt.Println("begin...")

	msg := &Msg{
		Name:	"QIAOYANG",
		Age:	33,
		Score:	99,
	}
	res, _ = json.Marshal(msg)


	http.HandleFunc("/", handler)
	http.ListenAndServe(":8081", nil)

	fmt.Println("exit...")
}


