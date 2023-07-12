package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("https://google.com")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	//interface method to write into terminal
	/*contents := make([]byte, 99999)
	resp.Body.Read(contents)
	fmt.Print(string(contents))
	resp.Body.Close()*/

	//interface method to write into terminal
	io.Copy(os.Stdout, resp.Body)

}
