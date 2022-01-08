package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {
	//resp, _ := http.Get("http://example.com")
	//defer resp.Body.Close()
	//body, _ := ioutil.ReadAll(resp.Body)
	//fmt.Println(string(body))
	base, _ := url.Parse("https://example.com/")
	reference, _ := url.Parse("/test?a=1&b=2")
	endpoint := base.ResolveReference(reference).String()
	fmt.Println(endpoint)
	// req, _ := http.NewRequest("GET", endpoint, nil)
	req, _ := http.NewRequest("POST", endpoint, bytes.NewBuffer([]byte("password")))
	req.Header.Add("If-None-Match", `W/"wyzzy"`)
	q := req.URL.Query()
	q.Add("c", "3&%")
	fmt.Println(q)
	fmt.Println(q.Encode())
	req.URL.RawQuery = q.Encode()

	var client *http.Client = &http.Client{}
	resp, _ := client.Do(req)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))

}
