package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"
)

var t = flag.String("t", "", "需要查询的电影名称")

func main() {
	flag.Parse()
	movie, err := GetMovieInfo(*t)

	if err != nil {
		fmt.Printf("get movie info failed: %s\n", err.Error())
		os.Exit(0)
	}
	fmt.Printf("The Poster of《%s》is %s\n", movie.Title, movie.Poster)
}

func GetMovieInfo(movieName string) (*Movie, error) {
	requestUrl := fmt.Sprintf(posterUrl, apiKey, movieName)

	resp, err := http.Get(requestUrl)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("Request Movie info failed: %s", resp.Status)
	}

	var movieInfo Movie

	if err = json.NewDecoder(resp.Body).Decode(&movieInfo); err != nil {
		resp.Body.Close()
		return nil, err
	}

	resp.Body.Close()
	return &movieInfo, nil
}
