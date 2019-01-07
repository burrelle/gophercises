package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	adventure "github.com/burrelle/gophercises/choose-your-own-adventure"
)

func main() {
	port := flag.Int("port", 3000, "the port to start the CYOA web application on")
	filename := flag.String("file", "gopher.json", "the JSON file with CYOA story")
	flag.Parse()
	fmt.Printf("Using the story in %s", *filename)

	f, err := os.Open(*filename)
	if err != nil {
		panic(err)
	}

	story, err := adventure.JsonStory(f)
	if err != nil {
		panic(err)
	}

	h := adventure.NewHandler(story, nil)
	fmt.Printf("\nStarting the server at: http://localhost:%d", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), h))
}
