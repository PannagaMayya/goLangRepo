package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

// Models
type Country struct {
	Countryid   string   `json`
	CountryName string   `json`
	States      []*State `json`
}

type State struct {
	Stateid   string `json`
	Statename string `json`
}

// Handler type
type apiHandler struct{}

func (apiHandler) ServeHTTP(a http.ResponseWriter, b *http.Request) {
	if b.Method == "GET" {
		a.Write([]byte(time.Now().Format(time.ANSIC)))
		fmt.Println("API Page Loaded")
	} else {
		notFound(a, b)
	}
}

// DB
var countries []Country

// Middleware
func (c *Country) IsValid() bool {
	data, _ := json.Marshal(c)
	return json.Valid(data)
}
func main() {

	mux := http.NewServeMux()
	//Routes
	mux.HandleFunc("/api/getAllCountries", getAllCountries)
	mux.HandleFunc("/api/createCountry", createCountry)
	mux.HandleFunc("/api/getOneCountry", getOneCountry)
	mux.Handle("/api", apiHandler{})

	//Route URL
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		// The "/" pattern matches everything, so we need to check
		// that we're at the root here.
		if req.URL.Path != "/" || req.Method == "POST" {
			notFound(w, req)
		}

		fmt.Fprintf(w, "Welcome to the home page!")
		fmt.Println("Home Page Loaded")
	})
	fmt.Println("Server is Running on Port 8080...")
	log.Fatal(http.ListenAndServe(":8080", mux))

}

// Controllers
func getAllCountries(w http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(countries)
	} else {
		notFound(w, req)
	}
}

func getOneCountry(w http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		w.Header().Set("Content-Type", "application/json")
		params := strings.Split(req.URL.RawQuery, "=")
		fmt.Println(params[1])
		for _, d := range countries {
			if d.Countryid == params[1] {
				json.NewEncoder(w).Encode(d)
				return
			}
		}
		json.NewEncoder(w).Encode("No Data Found")

	} else {
		notFound(w, req)
	}
}

func createCountry(w http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		w.Header().Set("Content-Type", "application/json")
		//_ = json.NewEncoder(w).Encode(countries)
		//Need a fix for below - req.Body is of type "http.noBody{}"
		if req.Body == nil {
			json.NewEncoder(w).Encode("No Data Entered")
			return
		}
		var country Country
		json.NewDecoder(req.Body).Decode(&country)
		for _, d := range countries {

			if d.Countryid == country.Countryid {
				json.NewEncoder(w).Encode("Id Already exist")
				return
			}
		}

		countries = append(countries, country)
		json.NewEncoder(w).Encode("Successful Update")
	} else {
		notFound(w, req)
	}
}

func notFound(w http.ResponseWriter, req *http.Request) {

	fmt.Println(req.URL.Path, "/", req.Method, "is Invalid URL/Invalid Method")
	http.NotFound(w, req)
	return

}
