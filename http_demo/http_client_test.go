package http_demo

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"testing"
)

func TestHttpGet(t *testing.T) {
	response, err := http.Get("https://blog.csdn.net/zyndev")
	if err != nil {
		t.Error(err)
		panic(err)
	}
	t.Log(response)
}

func TestHttpGetHeader(t *testing.T) {
	targetUrl := "https://blog.csdn.net/zyndev"

	req, _ := http.NewRequest("GET", targetUrl, nil)

	req.Header.Add("Authorization", "xxxx")

	response, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Error(err)
		panic(err)
	}
	t.Log(response)
}

func TestHttpPost(t *testing.T) {
	targetUrl := "https://blog.csdn.net/zyndev"

	payload := strings.NewReader("a=111")

	response, err := http.Post(targetUrl, "x-www-form-urlencoded", payload)

	if err != nil {
		t.Error(err)
		panic(err)
	}
	t.Log(response)
}

func TestHttpPostForm(t *testing.T) {
	targetUrl := "https://blog.csdn.net/zyndev"

	payload := url.Values{"key":{"value"}, "id": {"123"}}

	response, err := http.PostForm(targetUrl, payload)

	if err != nil {
		t.Error(err)
		panic(err)
	}
	t.Log(response)
}

func TestHttpPut(t *testing.T)  {
	targetUrl := "http://xxxxx:8080/v2/repos/wh_flowDataSource1"

	payload := strings.NewReader("{\n    \"schema\": [\n      {\n        \"key\": \"a\",\n        \"valtype\": \"string\",\n        \"required\": false\n      }\n    ]\n}")

	req, _ := http.NewRequest("PUT", targetUrl, payload)

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "bmgAAGI155F6MJ3N2Tk9ruL_6XQpx-uxkkg:yKx_OYDtI3njD7-c7Y87Oov0GpI=:eyJyZXNvdXJBvcy93aF9mbG93RGF0YVNvdXJjZTEiLCJleHBpcmVzIjoxNTM2NzU1MjkwLCJjb250ZW50TUQ1IjoiIiwiY29udGVudFR5cGUiOiJhcHBsaWNhdGlvbi9qc29uIiwiaGVhZGVycyI6IiIsIm1ldGhvZCI6IlBVVCJ9")
	req.Header.Add("Date", "Wed, 12 Sep 2018 02:10:09 GMT")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))
}
