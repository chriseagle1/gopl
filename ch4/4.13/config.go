package main

const apiKey = "3d53779c"
const posterUrl = "http://www.omdbapi.com/?apikey=%s&t=%s"

type Movie struct {
	Title string
	Year string
	Rated string
	Released string
	Runtime string
	Genre string
	Director string
	Writer string
	Actors string
	Plot string
	Language string
	Country string
	Awards string
	Poster string
	Ratings []*Rating
	Metascore string
	ImdbRating string
	ImdbVotes string
	ImdbID string
	Type string
	DVD string
	BoxOffice string
	Production string
	Website string
	Response string

}

type Rating struct {
	Source string
	Value string
}