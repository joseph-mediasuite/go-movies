package models

import (
	"time"
)

type Movie struct {
	Id          int       `json:"id"`
	Title       string    `json:"title"`
	ReleaseDate time.Time `json:"release_date"`
	Runtime     int       `json:"runtime"`
	MPAARating  string    `json:"mpaa_rating"`
	Description string    `json:"description"`
	Image       string    `json:"image"`
	CreatedAt   time.Time `json:"-"`
	UpdatedAt   time.Time `json:"-"`
	Genres []*Genre `json:"genres,omitempty"`
	GenresArray []int `json:"genres_array,omitempty"`

}

type Genre struct {
	Id int `json:"id"`
	Genre string `json:"genre"`
	Checked bool `json:"checked"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
	
}
