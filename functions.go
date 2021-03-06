package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// Unsplash structure
type Unsplash struct {
	ID   string `json:"id"`
	Desc string `json:"description"`
	Urls struct {
		Raw   string `json:"raw"`
		Small string `json:"small"`
	} `json:"urls"`
	Links struct {
		HTML string `json:"html"`
	} `json:"links"`
	User struct {
		Name string `json:"name"`
	} `json:"user"`
	Exif struct {
		Make  string `json:"make"`
		Model string `json:"model"`
	} `json:"exif"`
}

// SearchInfo structure
type SearchInfo struct {
	Photos struct {
		Results []Unsplash `json:"results"`
	} `json:"photos"`
}

func random() Unsplash {
	url := fmt.Sprintf("%sphotos/random?client_id=%s", APIURL, os.Getenv("CLIENT_ACCESS_KEY"))
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var info Unsplash
	err = json.Unmarshal(responseData, &info)
	if err != nil {
		fmt.Print(err)
	}
	return info
}

func search(q string) SearchInfo {
	url := fmt.Sprintf("%ssearch?query=%s/photos&client_id=%s", APIURL, q, os.Getenv("CLIENT_ACCESS_KEY"))
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	var info SearchInfo
	err = json.Unmarshal(responseData, &info)
	if err != nil {
		log.Fatal(err)
	}
	return info
}
