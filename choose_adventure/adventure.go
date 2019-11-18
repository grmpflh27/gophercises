package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func loadStory(storyPath string) map[string]interface{} {
	fp, err := os.Open(storyPath)
	if err != nil {
		log.Fatal("Could not open story json")
	}

	defer fp.Close()
	byteValue, _ := ioutil.ReadAll(fp)

	// var chapterMap map[string]chapter
	// json.Unmarshal(byteValue, &chapterMap)

	var result map[string]interface{}
	json.Unmarshal(byteValue, &result)

	return result
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "CHOOSE YOUR OWN ADVENTURE!")
	fmt.Println("Endpoint Hit: homePage")
}

func handleRequests(chapters map[string]interface{}) {
	myRouter := mux.NewRouter().StrictSlash(true)
	for k, _ := range chapters {
		fmt.Println(k)
		myRouter.HandleFunc("/"+k, func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "!!!!")
		})
	}

	myRouter.HandleFunc("/", home)
	http.ListenAndServe(":10000", myRouter)
}

func main() {
	fmt.Println("Starting this adventure")

	// 1) loading story
	storyJSONPath := "gopher.json"
	chapters := loadStory(storyJSONPath)

	// 2) generate html templates

	// 3) define and start webserver

	handleRequests(chapters)
}

// type Chapter struct {
// 	title   string           `json:"title"`
// 	story   []string         `json:"story"`
// 	options []ChapterOptions `json:"options"`
// }

// type ChapterOptions struct {
// 	text string `json:"text"`
// 	arc  string `json:"arc"`
// }
