package main

import "gorm.io/gorm"

type Movie struct {
	gorm.Model
	ID            uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Title         string `json:"title" gorm:"not null"`
	ReleaseDate   string `json:"release_date" gorm:"not null"`
	Runtime       string `json:"runtime" gorm:"not null"`
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
