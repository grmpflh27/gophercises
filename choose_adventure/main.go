package main

import (
	"fmt"
	"net/http"

	"github.com/grmpflh27/gophercises/choose_adventure/funkyadventure"
)

func main() {
	fmt.Println("Starting this adventure")
	storyJSONPath := "gopher.json"

	// loading story
	chapters := funkyadventure.LoadStory(storyJSONPath)

	// define routes and start webserver
	router := funkyadventure.SetupRouter(chapters)
	http.ListenAndServe(":10000", router)
}
