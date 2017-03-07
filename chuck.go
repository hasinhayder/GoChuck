package main

/* Display Random Chuck Norris Joke */

import (
	"fmt"

	"github.com/bndr/gopencils"
)

type Joke struct {
	T string `json:"type"`
	V struct {
		ID   int      `json:"id"`
		Joke string   `json:"joke"`
		Cats []string `json:"categories"`
	} `json:"value"`
}

func main() {
	randomJoke := &Joke{}
	url := "http://api.icndb.com/jokes/"
	api := gopencils.Api(url)
	api.Res("random", randomJoke).Get()
	fmt.Println(randomJoke.V.Joke)
}
