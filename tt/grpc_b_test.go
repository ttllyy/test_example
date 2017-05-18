package tt

import (
	"testing"
	"golang.org/x/net/context"
	//"fmt"
)

func BenchmarkGrpc(b *testing.B) {
	for i := 0; i < b.N; i++ { //use b.N for looping
		r, _ := Gclient.HandlerMsg(context.Background(), &rpcObj)
		r.Age++
		//fmt.Println(*r)
	}
}

