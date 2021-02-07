package http_demo

import (
	"fmt"
	"net/http"
	"testing"
)

func TestHttpGet(t *testing.T)  {
	response, err := http.Get("http://localhost:8010/actuator/health")
	if err != nil {
		t.Error(err)
		panic(err)
	}
	fmt.Println(response)
	if response.StatusCode == 200 {
		response.Body.
	}
	t.Log(response)
}
