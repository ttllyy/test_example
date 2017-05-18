package tt


import (
	"net/http"
	"encoding/json"
	"strings"
	pb "test_example/hmsg"
	"google.golang.org/grpc"
	"log"
)

type RestClient struct {
	Url        string       // The location of the server like "http://localhost:8065"
	HttpClient *http.Client // The http client
}

type Msg struct {
	Name	string		`json:"name"`
	Age		int			`json:"age"`
	Score	int			`json:"score"`
}


var Rclient *RestClient
var Gclient pb.HmsgClient
var str string
var rpcObj pb.Request

func NewClient() *RestClient {
	return &RestClient{"http://127.0.0.1:8081/", &http.Client{}}
}

func (c *RestClient) DoPost() (*http.Response, error) {
	rq, _ := http.NewRequest("POST", c.Url, strings.NewReader(str))
	return c.HttpClient.Do(rq)
}




func init() {
	Rclient = NewClient()
	msg := &Msg{
		Name:	"QIAOYANG",
		Age:	44,
		Score:	100,
	}
	b, _ := json.Marshal(msg)
	str = string(b)



	conn, err := grpc.Dial("127.0.0.1:8082", grpc.WithInsecure())
	if err != nil {
		log.Fatal("did not connect: %v", err)
	}

	Gclient = pb.NewHmsgClient(conn)

	rpcObj = pb.Request{
		Name: 	"qiaoyang",
		Age:	77,
		Score:	90,
	}
}

