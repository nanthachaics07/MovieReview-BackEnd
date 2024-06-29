package models

import (
	"gorm.io/gorm"
)

type Movies struct {
	gorm.Model
	ID            uint   `gorm:"primaryKey;autoIncrement" json:"id"`
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
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	ReleaseDate string `json:"release_date"`
	MPAA        string `json:"mpaa_rating"`
	ImageURL    string `json:"imageUrl"`
}

type Movie struct {
	gorm.Model
	ID            uint   `json:"id" gorm:"primaryKey"`
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

// type Movies struct {
// 	gorm.Model
// 	Title         string `gorm:"varchar(50);not null"`
// 	ReleaseDate   string `gorm:"varchar(20);not null"`
// 	Runtime       string `gorm:"varchar(20);not null"`
// 	Rating        string `gorm:"varchar(20);"`
// 	Category      string `gorm:"varchar(20);not null"`
// 	Popularity    string `gorm:"varchar(20);not null"`
// 	Budget        int    `gorm:"default:0;not null"`
// 	Revenue       int    `gorm:"default:0;not null"`
// 	Director      string `gorm:"varchar(50);not null"`
// 	Casting       string `gorm:"varchar(100);not null"`
// 	Writers       string `gorm:"varchar(50);not null"`
// 	DistributedBy string `gorm:"varchar(50);not null"`
// 	MPAA          string `gorm:"varchar(10);not null"`
// 	Description   string `gorm:"varchar(1000);not null"`
// 	ImageURL      string `gorm:"varchar(1000);not null"`
// }
