package models

import "gorm.io/gorm"

type Movies struct {
	gorm.Model
	Title         string `json:"title"`
	ReleaseDate   string `json:"release_date"`
	Runtime       string `json:"runtime"`
	Rating        string `json:"rating"`
	Category      string `json:"category"`
	Popularity    string `json:"popularity"`
	Budget        int    `json:"budget"`
	Revenue       int    `json:"revenue"`
	Director      string `json:"Director"`
	Casting       string `json:"casting"`
	Writers       string `json:"Writers"`
	DistributedBy string `json:"Distributed by"`
	MPAA          string `json:"mpaa_rating"`
	Description   string `json:"description"`
	ImageURL      string `json:"imageUrl"`
}

type MovieOnHomePage struct {
	Title       string `json:"title"`
	ReleaseDate string `json:"release_date"`
	MPAA        string `json:"mpaa"`
	ImageURL    string `json:"image_url"`
}
