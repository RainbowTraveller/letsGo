//Get content type of site
package main

import (
	"fmt"
	"net/http"
)

func getHeaderType(url string) error {
	resp, err := http.Get(url)
	if err != null {
		fmt.Printf("Error : %s\n", err)
		return
	}
	defer resp.Body.Close()
	ctype := resp.Header.Get("content-type")
	fmt.Println("%s -> %s \n", url, ctype)
}

func main() {

}
