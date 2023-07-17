package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	resp, err := http.Get("https://google.com")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	//get JSON
	data, _ := http.Get("https://api.github.com")
	dataJSON, _ := ioutil.ReadAll(data.Body)
	var out map[string]interface{}
	_ = json.Unmarshal(dataJSON, &out)
	fmt.Println(out["current_user_url"])
	//interface method to write into terminal
	/*contents := make([]byte, 99999)
	resp.Body.Read(contents)
	fmt.Print(string(contents))
	resp.Body.Close()*/

	//interface method to write into terminal
	io.Copy(os.Stdout, resp.Body)

	//Routine
	/*Pinging each Links 3 times with individual time interval of 5 seconds*/
	fmt.Println("")
	links := []string{"https://google.com", "https://flipkart.com", "https://github.com"}
	linkMap := map[string]int{links[0]: 0, links[1]: 0, links[2]: 0}
	c := make(chan string)
	//Flag for Forloop Exit
	flag := 0
	for _, link := range links {
		go getStatus(link, linkMap, c)
	}
	for chanelSlice := range c {
		if flag >= 2 {
			break
		}
		//Limit link hit to 3
		if linkMap[chanelSlice] <= 2 {
			go func(cs string) {
				time.Sleep(5 * time.Second)
				getStatus(cs, linkMap, c)
			}(chanelSlice)
		} else {
			flag += 1
			continue
		}
	}
}
func getStatus(l string, lm map[string]int, c chan string) {
	_, err := http.Get(l)

	if err != nil {
		fmt.Println(l, " is DOWN, Ping count:", lm[l]+1)
		lm[l] += 1
		c <- l
		return
	}
	fmt.Println(l, " is UP, Ping count:", lm[l]+1)
	lm[l] += 1
	c <- l
}
