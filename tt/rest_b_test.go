package tt


import (
	"testing"
	"io/ioutil"
	//"fmt"
	//"fmt"
	"encoding/json"
)

func BenchmarkRest(b *testing.B) {
	for i := 0; i < b.N; i++ { //use b.N for looping
		r, _ := Rclient.DoPost()
		body, _ := ioutil.ReadAll(r.Body)
		var m Msg
		json.Unmarshal(body, &m)
		//fmt.Println(string(body))
	}
}
