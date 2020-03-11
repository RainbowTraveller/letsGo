package main

import (
	"fmt"
	"net/http"
)

//func responseCleanup(resp http.Response) {
//}

func getContentType(url string) (string, error) {
	resp, err := http.Get(url)
	content := ""
	//defer responseCleanup(resp)
	if err != nil {
		fmt.Println("Can not get valid response")
	} else {
		content = resp.Header.Get("Content-type")
	}
	if resp != nil {
		resp.Body.Close()
	}
	return content, err
}

func main() {

	resp, err := getContentType("https://golang.org")
	if err == nil {
		fmt.Println(resp)
	} else {
		fmt.Println("Invalid response")
	}
}
