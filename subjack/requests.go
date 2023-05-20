package main

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"github.com/valyala/fasthttp"
	"time"
)

var (
	client = &fasthttp.Client{
		TLSConfig: &tls.Config{InsecureSkipVerify: true},
		MaxConnsPerHost: 1024, // Adjust the maximum number of connections per host as per your needs
	}
)

func main() {
	url := "example.com" // Replace with your desired URL
	ssl := true          // Replace with your desired SSL value
	timeout := 10        // Replace with your desired timeout in seconds

	body := get(url, ssl, timeout)
	fmt.Println(string(body))
}

func get(url string, ssl bool, timeout int) []byte {
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)

	req.SetRequestURI(site(url, ssl))
	req.Header.Add("Connection", "close")

	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	err := client.DoTimeout(req, resp, time.Duration(timeout) * time.Second)
	if err != nil {
		// Handle the error here
	}

	return resp.Body()
}

func site(url string, ssl bool) string {
	var site bytes.Buffer
	site.WriteString("http://")
	site.WriteString(url)

	if ssl {
		site.Reset()
		site.WriteString("https://")
		site.WriteString(url)
	}

	return site.String()
}
