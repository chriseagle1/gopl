package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Movie struct {
	Title string
	Year int `json:"released"`
	Color bool `json:"color,omitempty"`
	Actors []string
}

func main() {
	var movies  = []Movie{
		{Title:"Casablanca", Year:1942, Color:false, Actors:[]string{"Humphrey Bogart", "Ingrid Bergman"}},
		{Title:"Cool Hand Luke", Year:1967, Color:true, Actors:[]string{"Paul Newman"}},
		{Title:"Bullitt", Year:1968, Color:true, Actors:[]string{"Steve McQueen", "Jacqueline Bisset"}},
	}

	data, err := json.Marshal(movies)

	if err != nil {
		log.Fatalf("Json marshaling failed:%s", err)
	}

	fmt.Printf("%s\n", data)

	data, err = json.MarshalIndent(movies, "", "    ")

	if err != nil {
		log.Fatalf("Json marshaling failed:%s", err)
	}

	fmt.Printf("%s\n", data)

	var titles []struct{Title string}

	var actorList []struct{Actors []string}

	if err = json.Unmarshal(data, &titles); err != nil {
		log.Fatalf("Json unmarshaling failed:%s\n", err)
	}

	fmt.Println(titles)

	if err = json.Unmarshal(data, &actorList); err != nil {
		log.Fatalf("Json unmarshaling failed:%s\n", err)
	}

	fmt.Println(actorList)
}
