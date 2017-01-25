// JSON marshaling and unmarshaling
package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	type Movie struct {
		Title  string
		Year   int  `json:"released"`
		Color  bool `json:"color,omitempty"`
		Actors []string
	}
	var movies = []Movie {
		{Title: "Casablanca", Year: 1942, Color: false,
			Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
		{Title: "Cool Hand Luke", Year: 1967, Color: true,
			Actors: []string{"Paul Newman"}},
		{Title: "Bullitt", Year: 1968, Color: true,
			Actors: []string{"Steve McQueen", "Jacqueline Bisset"}},
		{Title: "Sholay", Year: 1975, Color: true,
			Actors: []string{"Dharmendra", "Amitabh Bachchan", "Amjad Khan", "Sanjeev Kumar"}},
	}

	enc, err := json.MarshalIndent(movies, "", "	") // indent; json.Marshal returns a bare byte slice
	if err != nil {
		fmt.Printf("Marshaling failed with %s\n", err)
		os.Exit(2)
	}
	fmt.Printf("%s\n", enc)
	//titles := make([]struct{ Title string }, 10)
	var titles []struct{Title string}
	json.Unmarshal(enc, &titles)
	fmt.Println(titles)
}
