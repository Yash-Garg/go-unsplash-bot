package main

import (
	"os"
)

// Unsplash structure
type Unsplash struct {
	ID   string `json:"id"`
	Desc string `json:"description"`
	Urls struct {
		Raw   string `json:"raw"`
		Thumb string `json:"thumb"`
	} `json:"urls"`
	Links struct {
		HTML string `json:"html"`
	} `json:"links"`
	User struct {
		Name string `json:"name"`
		Exif struct {
			Make  string `json:"make"`
			Model string `json:"model"`
		} `json:"exif"`
	} `json:"user"`
}

// UnsplashInfo struct
type UnsplashInfo struct {
	Info []Unsplash
}

func random() Unsplash {
	var AccessKey = os.Getenv("ACCESS_KEY")
	url := API_URL + "photos/random?client_id=" + AccessKey
	print(url)

	return Unsplash{}
}
