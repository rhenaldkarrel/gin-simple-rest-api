package main

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Bintang di Surga", Artist: "Noah", Price: 56.99},
	{ID: "2", Title: "Kampungan", Artist: "Slank", Price: 17.99},
	{ID: "3", Title: "Naif", Artist: "Naif", Price: 39.99},
}
